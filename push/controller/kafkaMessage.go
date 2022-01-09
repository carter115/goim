package controller

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"gmimo/common/e"
	"gmimo/common/log"
	messageProto "gmimo/common/proto/message"
	"gmimo/push/config"
	"time"
)

var (
	KafkaConsumerGroup sarama.ConsumerGroup
)

func InitConsumerGroup(cf config.ConfKafka) (err error) {
	conf := sarama.NewConfig()
	conf.Version = sarama.V2_0_0_0 //  version
	conf.Consumer.Return.Errors = cf.ReturnErrors
	KafkaConsumerGroup, err = sarama.NewConsumerGroup(cf.Brokers, cf.GroupId, conf)
	return
}

// 从kafka拉取
func KafkaPull() {
	ctx := context.Background()
	for {
		if err := KafkaConsumerGroup.Consume(ctx, config.Kafka.Topics, &consumerGroupHandler{}); err != nil {
			log.Error(e.GetMsg(e.ERR_KAFKA_PULL_MSG), err)
			time.Sleep(time.Second) // 避免出错信息过多
		}
	}
}

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Debugf("[Kafka Message] topic:%q partition:%d offset:%d value:%s",
			msg.Topic, msg.Partition, msg.Offset, msg.Value)
		sess.MarkMessage(msg, "")

		// 把消息提交给调度器，进行转存
		obj := messageProto.Request{}
		if err := json.Unmarshal(msg.Value, &obj); err == nil {
			Scheduler.Push(&obj)
		}
	}
	return nil
}
