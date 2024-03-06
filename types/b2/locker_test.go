package b2

import (
	"github.com/kpyramid/bitcoin-inscribe/types"
	"log"
	"testing"
)

func TestGetLockerAddress(t *testing.T) {
	svc := types.GetServiceContext()

	address, err := GetLockAddress(svc, 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("address: %s", address)
}
