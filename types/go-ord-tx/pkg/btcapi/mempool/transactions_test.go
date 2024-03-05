package mempool

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"minter/types"
	"testing"
)

func TestGetRawTransaction(t *testing.T) {
	svc := types.GetServiceContext()
	client := NewClient(svc.Config.MempoolAddress, &chaincfg.SigNetParams)
	txId, _ := chainhash.NewHashFromStr("b752d80e97196582fd02303f76b4b886c222070323fb7ccd425f6c89f5445f6c")
	transaction, err := client.GetRawTransaction(txId)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(transaction.TxHash().String())
	}
}
