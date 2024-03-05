package task

import (
	"bytes"
	"encoding/hex"
	"github.com/btcsuite/btcd/wire"
	"github.com/kpyramid/bitcoin-inscribe/types"
	"github.com/kpyramid/bitcoin-inscribe/types/schema"
	"github.com/kpyramid/bitcoin-inscribe/types/wrap"
)

func broadcastCommitTx(svc *types.ServiceContext, order *schema.InscribeOrder) error {
	tx := wire.NewMsgTx(wire.TxVersion)
	txBytes, err := hex.DecodeString(order.CommitData)
	if err != nil {
		return err
	}
	if err := tx.Deserialize(bytes.NewReader(txBytes)); err != nil {
		return err
	}

	txHash, err := svc.Client.SendRawTransaction(tx, false)
	if err != nil {
		return err
	}

	// update order
	result := svc.Db.Model(&schema.InscribeOrder{}).Where("id = ?", order.Model.ID).
		Where("status = ?", schema.OrderStatusCommitBroadcast).
		Where("commit_tx_hash = ?", txHash.String()).
		Update("status", schema.OrderStatusRevealBroadcast)
	if err := wrap.IsDBError(result); err != nil {
		return err
	}
	return nil
}

func broadcastRevealTx(svc *types.ServiceContext, order *schema.InscribeOrder) error {
	tx := wire.NewMsgTx(wire.TxVersion)
	txBytes, err := hex.DecodeString(order.RevealData)
	if err != nil {
		return err
	}
	if err := tx.Deserialize(bytes.NewReader(txBytes)); err != nil {
		return err
	}

	txHash, err := svc.Client.SendRawTransaction(tx, false)
	if err != nil {
		return err
	}

	// update order
	result := svc.Db.Model(&schema.InscribeOrder{}).Where("id = ?", order.Model.ID).
		Where("status = ?", schema.OrderStatusRevealBroadcast).
		Where("reveal_tx_hash = ?", txHash.String()).
		Update("status", schema.OrderStatusUnConfirmed)
	if err := wrap.IsDBError(result); err != nil {
		return err
	}
	return nil
}
