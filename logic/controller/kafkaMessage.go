package controller

import (
	"context"
	"github.com/Shopify/sarama"
	"gmimo/common/e"
	"gmimo/common/log"
	"gmimo/logic/config"
)

// 处理消息的业务
// 1. 推送到MQ，返回结果给调用方；
// 2. [Push服务]从MQ拿取消息，转存到Mongodb，并且推送消息给[Comet服务]

var (
	KafkaProducer sarama.SyncProducer
)

func InitKafkaProducer(cf config.ConfKafka) (err error) {
	if KafkaProducer == nil {
		conf := sarama.NewConfig()
		conf.Producer.RequiredAcks = sarama.WaitForAll // 等待服务器所有副本都保存成功后的响应
		conf.Producer.Retry.Max = cf.MaxProduceRetry
		conf.Producer.Return.Successes = cf.ReturnSuccesses // 是否等待成功和失败后的响应

		KafkaProducer, err = sarama.NewSyncProducer(cf.Brokers, conf)
		return
	}
	return nil
}

// 推送到Kafka
func KafkaPush(ctx context.Context, content []byte) (err error) {
	if KafkaProducer == nil {
		err = InitKafkaProducer(config.Kafka)
		log.Error("init kafka producer error:", err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic:     config.Kafka.Topic,
		Partition: config.Kafka.Partitions,
		Value:     sarama.ByteEncoder(content),
	}

	part, offset, err := KafkaProducer.SendMessage(msg)
	if err != nil {
		log.Error(e.GetMsg(e.ERR_KAFKA_PUSH_MSG), err)
		return
	}
	log.Debugcf(ctx, "[Kafka Message] topic:%q partition:%d offset:%d value:%s",
		msg.Topic, part, offset, msg.Value)
	return nil
}
