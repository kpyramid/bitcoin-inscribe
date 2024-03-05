package task

import (
	"github.com/kpyramid/bitcoin-inscribe/types"
	"github.com/kpyramid/bitcoin-inscribe/types/schema"
	"github.com/kpyramid/bitcoin-inscribe/types/wrap"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

const EnableCoroutine = false

func HandleDispatcher(svc *types.ServiceContext) {
	for {
		loopStartID := uint(0)
		for {
			var orders []schema.InscribeOrder
			if err := svc.Db.
				Where("status in ?", []string{schema.OrderStatusInscribing, schema.OrderStatusCommitBroadcast, schema.OrderStatusRevealBroadcast}).
				Where("id > ?", loopStartID).
				Limit(10).Find(&orders).Error; err != nil {
				log.Errorf("get orders bridge_confirmed failed.error: %s", err)
			}
			if len(orders) == 0 {
				break
			}

			waitGroup := sync.WaitGroup{}
			for _, order := range orders {
				if EnableCoroutine {
					waitGroup.Add(1)
					go dispatcher(svc, &order)
				} else {
					if err := dispatcher(svc, &order); err != nil {
						log.Errorf("handle order failed. error: %s", wrap.WithErrorf("error: %s", err))
					}
				}

				// @TIP sleep for limit third party api request frequency
				time.Sleep(time.Second)
				loopStartID = order.Model.ID
			}
			waitGroup.Wait()
		}

		time.Sleep(time.Second * 5)
	}
}

func dispatcher(svc *types.ServiceContext, order *schema.InscribeOrder) error {
	switch order.Status {
	case schema.OrderStatusInscribing:
		if err := inscribeOrder(svc, order); err != nil {
			return err
		}
	case schema.OrderStatusCommitBroadcast:
		if err := broadcastCommitTx(svc, order); err != nil {
			return err
		}
	case schema.OrderStatusRevealBroadcast:
		if err := broadcastRevealTx(svc, order); err != nil {
			return err
		}
	}
	return nil
}
