package service

import (
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/pkg/btcapi"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

type UnisatClient struct {
	baseURL string
	apikey  string
}

func NewUnisatClient(netParams *chaincfg.Params, apikey string) *UnisatClient {
	baseURL := ""
	if netParams.Net == wire.MainNet {
		baseURL = "https://open-api.unisat.io"
	} else if netParams.Net == wire.TestNet3 {
		baseURL = "https://open-api-testnet.unisat.io"
	} else {
		log.Fatal("mempool don't support other netParams")
	}
	return &UnisatClient{
		baseURL: baseURL,
		apikey:  apikey,
	}
}

func (u UnisatClient) ListUnspentNonInscription(address btcutil.Address) ([]*btcapi.UnspentOutput, error) {
	url := fmt.Sprintf("%s%s", u.baseURL, fmt.Sprintf("/v1/indexer/address/%s/utxo-data?size=2000", address.EncodeAddress()))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", u.apikey))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// parse json
	parseData := gjson.Parse(string(body))
	getResult := parseData.Get("data.utxo")
	if !getResult.Exists() || !getResult.IsArray() {
		return nil, fmt.Errorf("get unisat utxo result failed. body: %s", parseData.String())
	}

	outputs := make([]*btcapi.UnspentOutput, 0)
	for _, getUTXO := range getResult.Array() {
		txID := getUTXO.Get("txid")
		if !txID.Exists() {
			return nil, fmt.Errorf("get unisat utxo txid result failed. value: %v", txID)
		}
		satoshi := getUTXO.Get("satoshi")
		if !satoshi.Exists() {
			return nil, fmt.Errorf("get unisat utxo satoshi result failed. value: %v", satoshi)
		}
		vout := getUTXO.Get("vout")
		if !satoshi.Exists() {
			return nil, fmt.Errorf("get unisat utxo vout result failed. value: %v", vout)
		}

		txHash, err := chainhash.NewHashFromStr(txID.String())
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, &btcapi.UnspentOutput{
			Value:    satoshi.Int(),
			Outpoint: wire.NewOutPoint(txHash, uint32(vout.Int())),
			Output:   wire.NewTxOut(satoshi.Int(), address.ScriptAddress()),
		})
	}

	for i, output := range outputs {
		log.Infof("get btc address utxo. id: %d, hash: %s, value: %d", i, output.Outpoint.Hash, output.Value)
	}

	return outputs, nil
}

func (u UnisatClient) ListUnspentIsInscription(address btcutil.Address) ([]*btcapi.UnspentOutput, error) {
	url := fmt.Sprintf("%s%s", u.baseURL, fmt.Sprintf("/v1/indexer/address/%s/inscription-utxo-data?size=2000", address.EncodeAddress()))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", u.apikey))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// parse json
	parseData := gjson.Parse(string(body))
	getResult := parseData.Get("data.utxo")
	if !getResult.Exists() || !getResult.IsArray() {
		return nil, fmt.Errorf("get unisat utxo result failed. body: %s", parseData.String())
	}

	outputs := make([]*btcapi.UnspentOutput, 0)
	for _, getUTXO := range getResult.Array() {
		if getUTXO.Get("isSpent").Bool() {
			log.Infof("get utxo isSpent is true")
			continue
		}

		txID := getUTXO.Get("txid")
		if !txID.Exists() {
			return nil, fmt.Errorf("get unisat utxo txid result failed. value: %v", txID)
		}
		satoshi := getUTXO.Get("satoshi")
		if !satoshi.Exists() {
			return nil, fmt.Errorf("get unisat utxo satoshi result failed. value: %v", satoshi)
		}
		vout := getUTXO.Get("vout")
		if !satoshi.Exists() {
			return nil, fmt.Errorf("get unisat utxo vout result failed. value: %v", vout)
		}

		txHash, err := chainhash.NewHashFromStr(txID.String())
		if err != nil {
			return nil, err
		}

		// parse brc20
		var brc20Content []*btcapi.BRC20Context
		inscriptions := getUTXO.Get("inscriptions")
		if !inscriptions.IsArray() {
			return nil, fmt.Errorf("get unisat utxo inscriptions result failed. value: %v", inscriptions)
		}
		for _, result := range inscriptions.Array() {
			if result.Get("isBRC20").Bool() {
				brc20Content = append(brc20Content, &btcapi.BRC20Context{
					InscriptionNumber: result.Get("inscriptionNumber").Int(),
					InscriptionId:     result.Get("inscriptionId").String(),
					Offset:            result.Get("offset").Int(),
					Moved:             result.Get("moved").Bool(),
					Sequence:          result.Get("sequence").Int(),
				})
			}
		}

		outputs = append(outputs, &btcapi.UnspentOutput{
			Value:        satoshi.Int(),
			Outpoint:     wire.NewOutPoint(txHash, uint32(vout.Int())),
			Output:       wire.NewTxOut(satoshi.Int(), address.ScriptAddress()),
			BRC20Context: brc20Content,
		})
	}

	for i, output := range outputs {
		log.Infof("get inscription address utxo. id: %d, hash: %s, value: %d", i, output.Outpoint.Hash, output.Value)
	}

	return outputs, nil
}

type TransferableInscription struct {
	InscriptionNumber  int64
	InscriptionId      string
	Value              int64
	Tick               string
	TransferableAmount int64
}

func (u UnisatClient) ListTransferableInscriptions(address btcutil.Address, tick string) ([]*TransferableInscription, error) {
	url := fmt.Sprintf("%s%s", u.baseURL, fmt.Sprintf("/v1/indexer/address/%s/brc20/%s/transferable-inscriptions?limit=20", address.EncodeAddress(), tick))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", u.apikey))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	// parse json
	parseData := gjson.Parse(string(body))
	getResult := parseData.Get("data.detail")
	if !getResult.Exists() || !getResult.IsArray() {
		return nil, fmt.Errorf("get unisat utxo result failed. body: %s", parseData.String())
	}

	transferableInscription := make([]*TransferableInscription, 0)
	for _, getUTXO := range getResult.Array() {
		inscriptionNumber := getUTXO.Get("inscriptionNumber")
		if !inscriptionNumber.Exists() {
			return nil, fmt.Errorf("get unisat utxo inscriptionNumber result failed. value: %v", inscriptionNumber)
		}
		inscriptionId := getUTXO.Get("inscriptionId")
		if !inscriptionId.Exists() {
			return nil, fmt.Errorf("get unisat utxo inscriptionId result failed. value: %v", inscriptionId)
		}
		satoshi := getUTXO.Get("satoshi")
		if !satoshi.Exists() {
			return nil, fmt.Errorf("get unisat utxo satoshi result failed. value: %v", satoshi)
		}
		getTick := getUTXO.Get("data.tick")
		if !getTick.Exists() {
			return nil, fmt.Errorf("get unisat utxo tick result failed. value: %v", getTick)
		}
		transferableAmount := getUTXO.Get("data.amt")
		if !transferableAmount.Exists() {
			return nil, fmt.Errorf("get unisat utxo tick result failed. value: %v", transferableAmount)
		}

		transferableInscription = append(transferableInscription, &TransferableInscription{
			InscriptionNumber:  inscriptionNumber.Int(),
			InscriptionId:      inscriptionId.String(),
			Value:              satoshi.Int(),
			Tick:               getTick.String(),
			TransferableAmount: transferableAmount.Int(),
		})
	}

	for i, output := range transferableInscription {
		log.Infof("get transferable inscription. index: %d, tick: %s, amount: %d, id: %s", i, output.Tick, output.TransferableAmount, output.InscriptionId)
	}

	return transferableInscription, nil
}
