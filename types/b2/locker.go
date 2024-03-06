package b2

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/kpyramid/bitcoin-inscribe/types"
	"math/big"
)

func GetLockAddress(svc *types.ServiceContext, tokenId int64) (string, error) {
	contractAddress := common.HexToAddress(svc.Config.B2ContractAddress)

	// call contract
	contract, err := NewB2(contractAddress, svc.B2Client)
	if err != nil {
		return "", err
	}

	result, err := contract.TokenLockedAddress(nil, big.NewInt(tokenId))
	if err != nil {
		return "", err
	}
	return result.L1Address, nil
}
