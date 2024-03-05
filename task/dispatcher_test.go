package task

import (
	"github.com/kpyramid/bitcoin-inscribe/types"
	"testing"
)

func TestDispatcher(t *testing.T) {
	svc := types.GetServiceContext()

	HandleDispatcher(svc)
}
