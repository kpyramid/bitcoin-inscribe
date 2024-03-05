package rpc

import (
	"context"
	"github.com/kpyramid/bitcoin-inscribe/proto"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLaunchInscribe(t *testing.T) {
	i := Inscribe{}
	_, err := i.LaunchInscribe(context.TODO(), &proto.LaunchInscribeRequest{OrderId: 1})
	if err != nil {
		log.Fatal(err)
	}
}
