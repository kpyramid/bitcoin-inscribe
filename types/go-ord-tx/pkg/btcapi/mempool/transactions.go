package mempool

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
)

func (c *MempoolClient) GetRawTransaction(txHash *chainhash.Hash) (*wire.MsgTx, error) {
	res, err := c.request(http.MethodGet, fmt.Sprintf("/tx/%s/raw", txHash.String()), nil)
	if err != nil {
		return nil, err
	}

	tx := wire.NewMsgTx(wire.TxVersion)
	if err := tx.Deserialize(bytes.NewReader(res)); err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *MempoolClient) BroadcastTx(tx *wire.MsgTx) (*chainhash.Hash, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return nil, err
	}

	res, err := c.request(http.MethodPost, "/tx", strings.NewReader(hex.EncodeToString(buf.Bytes())))
	if err != nil {
		return nil, err
	}

	txHash, err := chainhash.NewHashFromStr(string(res))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to parse tx hash, %s", string(res)))
	}
	return txHash, nil
}

type FeeRateEstimate struct {
	FastestFee  int64
	HalfHourFee int64
	HourFee     int64
	EconomyFee  int64
	MinimumFee  int64
}

func (c *MempoolClient) GetFeeRateEstimate() (*FeeRateEstimate, error) {
	resp, err := http.Get("https://mempool.space/api/v1/fees/recommended")
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("get fee rate statsu: %d", resp.StatusCode)
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	feeRateEstimate := FeeRateEstimate{}
	if err := json.Unmarshal(res, &feeRateEstimate); err != nil {
		return nil, err
	}
	return &feeRateEstimate, nil
}
