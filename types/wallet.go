package types

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/tyler-smith/go-bip39"
)

type HDWallet struct {
	rootKey *hdkeychain.ExtendedKey
}

func (hw *HDWallet) Init(netParams *chaincfg.Params, mnemonic string) error {
	rootKey, err := hdkeychain.NewMaster(bip39.NewSeed(mnemonic, ""), netParams)
	if err != nil {
		return err
	}
	hw.rootKey = rootKey
	return nil
}

func (hw HDWallet) Generate(n uint32) (string, error) {
	rootKey := hw.rootKey
	childKey, _ := rootKey.Derive(hdkeychain.HardenedKeyStart + 86) // 86'
	childKey, _ = childKey.Derive(hdkeychain.HardenedKeyStart + 0)  // 0'
	childKey, _ = childKey.Derive(hdkeychain.HardenedKeyStart + 0)  // 0'
	childKey, _ = childKey.Derive(0)                                // 1
	senderKey, _ := childKey.Derive(n)                              // 0
	senderPrivateKey, _ := senderKey.ECPrivKey()
	privateKeyHex := hex.EncodeToString(senderPrivateKey.Serialize())
	return privateKeyHex, nil
}
