package mempool

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/kpyramid/bitcoin-inscribe/types/go-ord-tx/pkg/btcapi"
	"io"
)

type MempoolClient struct {
	baseURL string
}

func NewClient(mempoolAddress string, netParams *chaincfg.Params) *MempoolClient {
	baseURL := mempoolAddress
	// if netParams.Net == wire.MainNet {
	// 	baseURL = "https://blockstream.info/api"
	// } else if netParams.Net == wire.TestNet3 {
	// 	baseURL = "https://blockstream.info/testnet/api"
	// } else if netParams.Net == chaincfg.SigNetParams.Net {
	// 	baseURL = "https://blockstream.info/signet/api"
	// } else {
	// 	log.Fatal("mempool don't support other netParams")
	// }
	return &MempoolClient{
		baseURL: baseURL,
	}
}

func (c *MempoolClient) request(method, subPath string, requestBody io.Reader) ([]byte, error) {
	return btcapi.Request(method, c.baseURL, subPath, requestBody)
}

var _ btcapi.BTCAPIClient = (*MempoolClient)(nil)
