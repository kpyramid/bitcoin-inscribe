package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/kpyramid/bitcoin-inscribe/types/wrap"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

const sizeLimit = 4 * 1024 * 1024 // 4MB
const transactionConfirmedCount = uint64(2)
const TxSequenceFinal = 0xffffffff

func GetPrivateKey(privateHex string) (*btcec.PrivateKey, *btcec.PublicKey, error) {
	privateKeyBytes, err := hex.DecodeString(privateHex)
	if err != nil {
		return nil, nil, err
	}
	privateKey, _pubKey := btcec.PrivKeyFromBytes(privateKeyBytes)
	return privateKey, _pubKey, nil
}

func GetPublicAddress(publicKey *btcec.PublicKey, cfg *chaincfg.Params) (btcutil.Address, error) {
	taprootAddress, err := btcutil.NewAddressTaproot(schnorr.SerializePubKey(txscript.ComputeTaprootKeyNoScript(publicKey)), cfg)
	if err != nil {
		return nil, err
	}
	return taprootAddress, nil
}

func GetFileContent(url string) ([]byte, error) {
	// Send a HEAD request to get the Content-Length
	log.Infof("get resource url: %s", url)
	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, wrap.WithErrorf("get resource failed. error: %s", err)
	}
	defer resp.Body.Close()

	// Parse Content-Length header
	contentLength, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return nil, wrap.WithErrorf("get resource content length failed. error: %s", err)
	}

	// Check if file size is within the limit
	if contentLength > sizeLimit {
		return nil, fmt.Errorf("file is larger than 4MB")
	}

	// Download the file
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the file content
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func WaitQuitSignal(svc *ServiceContext) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	log.Infof("waiting process exited...")
	svc.QuitMutex.Lock()
	log.Infof("exited.")
}

func GenerateOutPoint(txHash string, index uint32) (*wire.OutPoint, error) {
	txid, err := chainhash.NewHashFromStr(txHash)
	if err != nil {
		return nil, wrap.WithErrorf("invalid hash. hash: %s", err)
	}

	return &wire.OutPoint{
		Hash:  *txid,
		Index: index,
	}, nil
}

func ValidateReceiptUTXO(svc *ServiceContext, outPoint *wire.OutPoint, estimatedFeeRate int64) error {
	rawTx, err := svc.Client.GetRawTransactionVerbose(&outPoint.Hash)
	if err != nil {
		return wrap.WithErrorf("get raw tx failed. error: %s", err)
	}

	// confirmed count
	if rawTx.Confirmations >= transactionConfirmedCount {
		log.Infof("validate utxo successfu. confirmations: %d, need confirmed count: %d", rawTx.Confirmations, transactionConfirmedCount)
		return nil
	}

	// feeRate
	totalInputAmount := decimal.Zero
	for _, vin := range rawTx.Vin {
		prevTxHash, _ := chainhash.NewHashFromStr(vin.Txid)
		prevTx, _ := svc.Client.GetRawTransactionVerbose(prevTxHash)
		btcValue := decimal.NewFromFloat(prevTx.Vout[vin.Vout].Value)
		totalInputAmount = totalInputAmount.Add(btcValue.Shift(8))
	}

	totalOutputAmount := decimal.Zero
	for _, vout := range rawTx.Vout {
		btcValue := decimal.NewFromFloat(vout.Value)
		totalOutputAmount = totalOutputAmount.Add(btcValue.Shift(8))
	}

	fee := totalInputAmount.Sub(totalOutputAmount)
	txVSize := rawTx.Vsize
	feeRate := fee.Div(decimal.NewFromInt(int64(txVSize))).Ceil().IntPart()
	log.Infof("fee: %s, txVSize: %d, feeRate: %d", fee.String(), txVSize, feeRate)

	passEstimatedFeeRate := int64(float64(estimatedFeeRate) * 0.8)
	if feeRate < passEstimatedFeeRate {
		return wrap.WithErrorf("tx feeRate too small to estimate fee rate. tx feeRate: %d, passEstimatedFeeRate: %d", feeRate, passEstimatedFeeRate)
	}

	// nSequence
	for i, vin := range rawTx.Vin {
		if vin.Sequence != 0xffffffff {
			return wrap.WithErrorf("tx vin sequence is replaceable transactions. index: %d, txId: %s", i, vin.Txid)
		}
	}

	return nil
}

func GetTxOut(svc *ServiceContext, outpoint *wire.OutPoint) (int64, error) {
	out, err := svc.Client.GetTxOut(&outpoint.Hash, outpoint.Index, true)
	if err != nil {
		return -1, nil
	}

	return decimal.NewFromFloat(out.Value).Shift(8).IntPart(), nil
}

func GetTxHex(tx *wire.MsgTx) (string, error) {
	var buf bytes.Buffer
	if err := tx.Serialize(&buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}
