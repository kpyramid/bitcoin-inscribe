package task

import (
	"github.com/kpyramid/bitcoin-inscribe/types"
	"github.com/kpyramid/bitcoin-inscribe/types/schema"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestInscribe(t *testing.T) {
	svc := types.GetServiceContext()

	order := schema.InscribeOrder{}
	if err := svc.Db.Where("status = ?", schema.OrderStatusInscribing).First(&order).Error; err != nil {
		log.Fatal(err)
	}

	if err := inscribeOrder(svc, &order); err != nil {
		log.Fatal(err)
	}
}
