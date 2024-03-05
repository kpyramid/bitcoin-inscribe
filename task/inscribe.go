package task

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/kpyramid/bitcoin-inscribe/types"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/ord"
	"github.com/kpyramid/bitcoin-inscribe/types/schema"
	"github.com/kpyramid/bitcoin-inscribe/types/wrap"
	log "github.com/sirupsen/logrus"
)

func inscribeOrder(svc *types.ServiceContext, order *schema.InscribeOrder) error {
	svc.QuitMutex.Lock()
	defer svc.QuitMutex.Unlock()

	// address
	orderPrivateKeyHex, err := svc.Wallet.Generate(uint32(order.ReceiptAddressNumber))
	if err != nil {
		return err
	}
	orderPrivateKey, orderPublicKey, err := types.GetPrivateKey(orderPrivateKeyHex)
	if err != nil {
		return err
	}
	orderAddress, err := types.GetPublicAddress(orderPublicKey, svc.NetParams)
	if err != nil {
		return err
	}

	// inscribe
	nftContent, err := pickNFTTokenID(svc, order.TokenId)
	if err != nil {
		return err
	}

	// inscribe
	{
		contentType := "text/html;charset=utf-8"
		unspentList, err := svc.UnisatClient.ListUnspentNonInscription(orderAddress)
		if err != nil {
			return wrap.WithErrorf("get inscribe utxo failed. error: %s, address: %s", err, orderAddress)
		}
		log.Infof("get unspent address: %s, list: %v", orderAddress, unspentList)

		commitTxOutPointList := make([]*wire.OutPoint, 0)
		commitTxPrivateKeyList := make([]*btcec.PrivateKey, 0)
		for i := range unspentList {
			commitTxOutPointList = append(commitTxOutPointList, unspentList[i].Outpoint)
			commitTxPrivateKeyList = append(commitTxPrivateKeyList, orderPrivateKey)
		}
		if len(commitTxOutPointList) == 0 {
			return fmt.Errorf("commit tx outPoint list is empty. address: %s", orderAddress)
		}

		dataList := make([]ord.InscriptionData, 0)
		dataList = append(dataList, ord.InscriptionData{
			ContentType: contentType,
			Body:        []byte(nftContent),
			Destination: order.UserAddress,
		})

		// change pkScript
		changeAddress, err := btcutil.DecodeAddress(order.UserAddress, svc.NetParams)
		if err != nil {
			return err
		}
		changePkScript, err := txscript.PayToAddrScript(changeAddress)
		if err != nil {
			return err
		}

		request := ord.InscriptionRequest{
			CommitTxOutPointList:   commitTxOutPointList,
			CommitTxPrivateKeyList: commitTxPrivateKeyList,
			CommitFeeRate:          order.FeeRate,
			FeeRate:                order.FeeRate,
			DataList:               dataList,
			TransferList:           []ord.TransferData{},
			ChangePkScript:         changePkScript,
			SingleRevealTxOnly:     false,
		}

		tool, err := ord.NewInscriptionToolWithBtcApiClient(svc, svc.NetParams, svc.BtcApiClient, &request)
		if err != nil {
			return err
		}
		commitTxHex, commitTxHash, err := tool.GetCommitTxHex()
		if err != nil {
			return err
		}
		log.Printf("commitTxHex %s, commitTxHash: %s", commitTxHex, commitTxHash)

		revealTxHexList, revealTxHashList, err := tool.GetRevealTxHexList()
		if err != nil {
			return fmt.Errorf("get reveal tx hex err, %v", err)
		}
		for i, revealTxHex := range revealTxHexList {
			log.Printf("revealTxHex %d %s, revealTxHash: %s", i, revealTxHex, revealTxHashList[i])
		}

		// update order
		result := svc.Db.Model(&schema.InscribeOrder{}).
			Where("id = ?", order.Model.ID).
			Where("status = ?", schema.OrderStatusInscribing).
			Updates(map[string]interface{}{
				"commit_data":    commitTxHex,
				"commit_tx_hash": commitTxHash,
				"RevealData":     revealTxHexList[0],
				"reveal_tx_hash": revealTxHashList[0],
				"status":         schema.OrderStatusCommitBroadcast,
			})
		if err := wrap.IsDBError(result); err != nil {
			return err
		}
	}

	log.Infof("inscribe success. order id: %d, order receipt address: %s", order.Model.ID, orderAddress)
	return nil
}

func pickNFTTokenID(svc *types.ServiceContext, tokenId int64) (string, error) {
	return "hello", nil
}
