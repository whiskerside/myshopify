package job

import (
	"fmt"

	workers "github.com/digitalocean/go-workers2"
	"github.com/thoas/go-funk"
	"github.com/whiskerside/myshopify/conf"
	"github.com/whiskerside/myshopify/consts"
	"github.com/whiskerside/myshopify/log"
	"github.com/whiskerside/myshopify/params"
)

var (
	producer *workers.Producer
	logs     = log.Logger()
)

func init() {
	producer, _ = workers.NewProducer(workers.Options{
		ServerAddr: fmt.Sprintf("%s:%d", conf.Env.Redis.Host, conf.Env.Redis.Port),
		Database:   conf.Env.Redis.Db,
		PoolSize:   conf.Env.Redis.PoolSize,
		ProcessID:  "1",
	})
	logs.Info().Msg("Init producer...")
}

func Enqueue(topic string, msgBody params.MsgBody) (string, error) {
	queueName := consts.QueueMiscs
	if funk.Contains(consts.QueueForProducts, topic) {
		queueName = consts.QueueProducts
	} else if funk.Contains(consts.QueueForOrders, topic) {
		queueName = consts.QueueOrders
	}
	return producer.Enqueue(queueName, topic, msgBody)
}
