package kafka

import (
	"github.com/Shopify/sarama"
	"time"
)

func GetConfig(cfg *Config) *sarama.Config {
	config := sarama.NewConfig()
	config.ClientID = cfg.ClientID
	config.Version = sarama.V3_2_0_0                          // kafka server的版本号
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 写到随机分区中，默认设置8个分区
	config.Producer.Return.Successes = true                   // sync必须设置这个
	config.Producer.RequiredAcks = sarama.WaitForAll          // 也就是等待foolower同步，才会返回
	config.Producer.Timeout = 5 * time.Second                 // 设置超时时间
	config.Producer.Return.Errors = true
	config.Consumer.Return.Errors = true
	config.Metadata.Full = false                                           // 不用拉取全部的信息
	config.Consumer.Offsets.AutoCommit.Enable = true                       // 自动提交偏移量，默认开启，说时候，我没找到手动提交。
	config.Consumer.Offsets.AutoCommit.Interval = time.Second              // 这个看业务需求，commit提交频率，不然容易down机后造成重复消费。
	config.Consumer.Offsets.Initial = sarama.OffsetOldest                  // 从最开始的地方消费，业务中看有没有需求，新业务重跑topic。
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // rb策略，默认就是range
	return config
}

func ProducerConnect(host []string, config *sarama.Config) sarama.SyncProducer {
	client, err := sarama.NewSyncProducer(host, config)
	if err != nil {
		panic("生产端连接失败")
	}
	return client
}

func ConsumerGroupConnect(host []string, group string, config *sarama.Config) sarama.ConsumerGroup {
	client, err := sarama.NewConsumerGroup(host, group, config)
	if err != nil {
		panic(err)
	}
	return client
}
