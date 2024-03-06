package rpc

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/google/uuid"
	"github.com/kpyramid/bitcoin-inscribe/proto"
	"github.com/kpyramid/bitcoin-inscribe/types"
	"github.com/kpyramid/bitcoin-inscribe/types/b2"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/ord"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/pkg/btcapi"
	"github.com/kpyramid/bitcoin-inscribe/types/schema"
	"github.com/kpyramid/bitcoin-inscribe/types/wrap"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Inscribe struct {
	proto.UnimplementedInscribeServiceServer
}

func (Inscribe) GenerateInscribeNFT(ctx context.Context, request *proto.GenerateInscribeNFTRequest) (*proto.GenerateInscribeNFTResponse, error) {
	svc := types.GetServiceContext()

	inscribeOrder := schema.InscribeOrder{
		Status: schema.OrderStatusPending,
	}

	// validate exist order
	if err := svc.Db.Where("token_id = ?", request.TokenId).First(&inscribeOrder).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	switch inscribeOrder.Status {
	case schema.OrderStatusFailure:
		fallthrough
	case schema.OrderStatusClosed:
		inscribeOrder.Status = schema.OrderStatusPending
	}
	if inscribeOrder.Status != schema.OrderStatusPending {
		return nil, fmt.Errorf("inscribe order in process")
	}

	if inscribeOrder.ReceiptAddressNumber == 0 {
		// create order wallet
		hdNumber, err := svc.Redis.Incr(context.TODO(), types.OrderWalletRedisKey).Result()
		if err != nil {
			return nil, err
		}
		inscribeOrder.ReceiptAddressNumber = hdNumber
	}

	orderPrivateKeyHex, err := svc.Wallet.Generate(uint32(inscribeOrder.ReceiptAddressNumber))
	if err != nil {
		return nil, err
	}
	_, orderPublicKey, err := types.GetPrivateKey(orderPrivateKeyHex)
	if err != nil {
		return nil, err
	}
	publicKey, err := types.GetPublicAddress(orderPublicKey, svc.NetParams)
	if err != nil {
		return nil, err
	}
	userAddressStr, err := b2.GetLockAddress(svc, request.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get btc mainnet locker address failed. error: %s", err)
	}

	// calc fee
	feeRate := request.FeeRate
	if feeRate < 10 {
		return nil, fmt.Errorf("fee rate can not too small. fee_rate: %d", feeRate)
	}

	// user address
	userAddress, err := btcutil.DecodeAddress(userAddressStr, svc.NetParams)
	if err != nil {
		return nil, err
	}

	_, totalAmount, err := estimateNFTTotalAmount(svc, userAddress, feeRate)
	if err != nil {
		return nil, err
	}

	// save db
	inscribeOrder.OrderId = uuid.New().String()
	inscribeOrder.UserAddress = userAddress.EncodeAddress()
	inscribeOrder.ReceiptAddress = publicKey.EncodeAddress()
	inscribeOrder.TotalAmount = totalAmount
	inscribeOrder.FeeRate = request.FeeRate
	inscribeOrder.NetWork = svc.NetParams.Name
	inscribeOrder.TokenId = request.TokenId
	inscribeOrder.Status = schema.OrderStatusPending
	if err := svc.Db.Save(&inscribeOrder).Error; err != nil {
		return nil, err
	}

	// estimate total amount
	resp := &proto.GenerateInscribeNFTResponse{
		ReceiptAddress: publicKey.EncodeAddress(),
		TotalAmount:    totalAmount,
		FeeRate:        feeRate,
		Network:        svc.NetParams.Name,
	}

	log.Infof("generate order wallet. address: %s, number: %d, amount: %d, fee_rate: %d", publicKey.EncodeAddress(), inscribeOrder.ReceiptAddressNumber, totalAmount, feeRate)
	return resp, nil
}

func (Inscribe) LaunchInscribe(ctx context.Context, request *proto.LaunchInscribeRequest) (*proto.LaunchInscribeResponse, error) {
	svc := types.GetServiceContext()
	order := &schema.InscribeOrder{}
	if err := svc.Db.Where("token_id = ?", request.TokenId).
		Where("status = ?", schema.OrderStatusPending).
		First(order).Error; err != nil {
		return nil, err
	}

	orderPrivateKeyHex, err := svc.Wallet.Generate(uint32(order.ReceiptAddressNumber))
	if err != nil {
		return nil, err
	}
	_, orderPublicKey, err := types.GetPrivateKey(orderPrivateKeyHex)
	if err != nil {
		return nil, err
	}
	orderAddress, err := types.GetPublicAddress(orderPublicKey, svc.NetParams)
	if err != nil {
		return nil, err
	}

	// get utxo
	utxo, err := svc.UnisatClient.ListUnspentNonInscription(orderAddress)
	if err != nil {
		log.Errorf("get non-inscribe utxo failed. error: %s", err)
		return nil, fmt.Errorf("get address utxo failed. address: %s, error: %s", orderAddress.EncodeAddress(), err)
	}
	var payAmount *btcapi.UnspentOutput = nil
	for _, u := range utxo {
		if u.Value >= order.TotalAmount {
			payAmount = u
		}
	}
	if payAmount == nil {
		log.Errorf("payAmount is nil, orderAddress: %s, utxo: %v", orderAddress.EncodeAddress(), utxo)
		return nil, fmt.Errorf("get balance failed. address: %s, orderTotalAmount: %d, utxo: %v", orderAddress.EncodeAddress(), order.TotalAmount, utxo)
	}

	// validate utxo
	if order.FeeRate <= svc.Config.MinFeeRate {
		log.Errorf("order amount is zero; order: %v", order)
		return nil, fmt.Errorf("fee rate too small. fee_rate: %d", order.FeeRate)
	}
	if err := types.ValidateReceiptUTXO(svc, payAmount.Outpoint, order.FeeRate); err != nil {
		return nil, err
	}

	// update status
	result := svc.Db.Model(&schema.InscribeOrder{}).Where("id = ?", order.Model.ID).
		Where("status = ?", schema.OrderStatusPending).
		Update("status", schema.OrderStatusInscribing)
	if err := wrap.IsDBError(result); err != nil {
		return nil, err
	}

	resp := &proto.LaunchInscribeResponse{}
	return resp, nil
}

func (Inscribe) GetInscribeInfo(ctx context.Context, request *proto.GetInscribeInfoRequest) (*proto.GetInscribeInfoResponse, error) {
	svc := types.GetServiceContext()

	inscribeOrder := schema.InscribeOrder{}
	if err := svc.Db.Where("token_id = ?", request.TokenId).First(&inscribeOrder).Error; err != nil {
		return nil, err
	}

	resp := proto.GetInscribeInfoResponse{
		TokenId:       inscribeOrder.TokenId,
		InscriptionId: inscribeOrder.InscriptionId,
		CommitTxHash:  inscribeOrder.CommitTxHash,
		RevealTxHash:  inscribeOrder.RevealTxHash,
		Status:        inscribeOrder.Status,
	}
	return &resp, nil
}

func estimateNFTTotalAmount(svc *types.ServiceContext, userAddress btcutil.Address, feeRate int64) (int64, int64, error) {
	var totalAmount int64 = 0
	var fee int64 = 0
	var changeMinimumSats int64 = 800

	// reveal + witness
	var nftTxVirtualSize int64 = 407

	var commitSize, revealSize int64
	switch userAddress.(type) {
	case *btcutil.AddressWitnessPubKeyHash:
		// P2WPKH
		{
			txCount := 1
			if changeMinimumSats != 0 {
				txCount += 1
			}
			switch txCount {
			case 1:
				return 0, 0, fmt.Errorf("not support")
			case 2:
				commitSize = 154
				revealSize = 82
			case 3:
				commitSize = 185
				revealSize = 82
			default:
				return 0, 0, fmt.Errorf("not support")
			}
		}
	case *btcutil.AddressTaproot:
		// P2TR
		{
			txCount := 1
			if changeMinimumSats != 0 {
				txCount += 1
			}
			switch txCount {
			case 1:
				return 0, 0, fmt.Errorf("not support")
			case 2:
				commitSize = 154
				revealSize = 94
			case 3:
				commitSize = 197
				revealSize = 94
			default:
				return 0, 0, fmt.Errorf("not support")
			}
		}
	case *btcutil.AddressPubKeyHash:
		// P2PKH
		{
			txCount := 1
			if changeMinimumSats != 0 {
				txCount += 1
			}
			switch txCount {
			case 1:
				return 0, 0, fmt.Errorf("not support")
			case 2:
				commitSize = 154
				revealSize = 85
			case 3:
				commitSize = 188
				revealSize = 85
			default:
				return 0, 0, fmt.Errorf("not support")
			}
		}
	case *btcutil.AddressScriptHash:
		// P2SH-P2PKH
		{
			txCount := 1
			if changeMinimumSats != 0 {
				txCount += 1
			}
			switch txCount {
			case 1:
				return 0, 0, fmt.Errorf("not support")
			case 2:
				commitSize = 154
				revealSize = 83
			case 3:
				commitSize = 186
				revealSize = 83
			default:
				return 0, 0, fmt.Errorf("not support")
			}
		}
	default:
		return 0, 0, wrap.WithErrorf("invalid address type. address: %s", userAddress)
	}

	fee = (commitSize+revealSize+nftTxVirtualSize)*feeRate + ord.DefaultRevealOutValue + changeMinimumSats
	totalAmount = fee

	return fee, totalAmount, nil
}
