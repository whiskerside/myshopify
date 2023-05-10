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
)

func init() {
	producer, _ = workers.NewProducer(workers.Options{
		ServerAddr: fmt.Sprintf("%s:%d", conf.Env.RedisHost, conf.Env.RedisPort),
		Database:   conf.Env.RedisDb,
		PoolSize:   conf.Env.RedisPoolSize,
		ProcessID:  "1",
	})
	log.Logger.Info().Msg("Init producer...")
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
