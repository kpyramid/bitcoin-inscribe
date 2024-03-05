package rpc

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/kpyramid/bitcoin-inscribe/proto"
	"github.com/kpyramid/bitcoin-inscribe/types"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/ord"
	"github.com/kpyramid/bitcoin-inscribe/types/schema"
	"github.com/kpyramid/bitcoin-inscribe/types/wrap"
	log "github.com/sirupsen/logrus"
)

type Inscribe struct {
	proto.UnimplementedInscribeServiceServer
}

func (Inscribe) GenerateInscribeNFTParams(ctx context.Context, request *proto.GenerateInscribeNFTParamsRequest) (*proto.GenerateInscribeNFTParamsResponse, error) {
	svc := types.GetServiceContext()

	// create order account
	hdNumber, err := svc.Redis.Incr(context.TODO(), types.OrderWalletRedisKey).Result()
	if err != nil {
		return nil, err
	}

	orderPrivateKeyHex, err := svc.Wallet.Generate(uint32(hdNumber))
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

	// calc fee
	feeRate := request.FeeRate
	if feeRate < 10 {
		return nil, fmt.Errorf("fee rate can not too small. fee_rate: %d", feeRate)
	}

	// user address
	userAddress, err := btcutil.DecodeAddress(request.UserAddress, svc.NetParams)
	if err != nil {
		return nil, err
	}

	_, totalAmount, err := estimateNFTTotalAmount(svc, userAddress, feeRate)
	if err != nil {
		return nil, err
	}

	// save db
	inscribeOrder := schema.InscribeOrder{
		OrderId:              request.OrderId,
		UserAddress:          request.UserAddress,
		ReceiptAddress:       publicKey.EncodeAddress(),
		ReceiptAddressNumber: hdNumber,
		FeeRate:              request.FeeRate,
		NetWork:              svc.NetParams.Name,
		TokenId:              request.TokenId,
		Status:               schema.OrderStatusPending,
	}
	if err := svc.Db.Create(&inscribeOrder).Error; err != nil {
		return nil, err
	}

	// estimate total amount
	resp := &proto.GenerateInscribeNFTParamsResponse{
		ReceiptAddress: publicKey.EncodeAddress(),
		TotalAmount:    totalAmount,
		FeeRate:        feeRate,
		Network:        svc.NetParams.Name,
	}

	log.Infof("generate order wallet. address: %s, number: %d, amount: %d, fee_rate: %d", publicKey.EncodeAddress(), hdNumber, totalAmount, feeRate)
	return resp, nil
}

func (Inscribe) LaunchInscribe(context.Context, *proto.LaunchInscribeRequest) (*proto.LaunchInscribeResponse, error) {
	panic("Todo")
}

func (Inscribe) GetInscribeInfo(context.Context, *proto.GetInscribeInfoRequest) (*proto.GetInscribeInfoResponse, error) {
	panic("Todo")
}

func estimateNFTTotalAmount(svc *types.ServiceContext, userAddress btcutil.Address, feeRate int64) (int64, int64, error) {
	var totalAmount int64 = 0
	var fee int64 = 0
	var changeMinimumSats int64 = 800

	// reveal + witness
	var nftTxVirtualSize int64 = 72

	var commitSize, revealSize int64
	switch userAddress.(type) {
	case *btcutil.AddressWitnessPubKeyHash:
		// P2WPKH
		{
			// commit + reveal(tx + witness[max nft])
			// const RevealOutValue= 546
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
			// commit + reveal(tx + witness[max nft])
			// const RevealOutValue= 546
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
			// commit + reveal(tx + witness[max nft])
			// const RevealOutValue= 546
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
			// commit + reveal(tx + witness[max nft])
			// const RevealOutValue= 546
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
