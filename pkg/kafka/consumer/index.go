package consumer

//import (
//	"context"
//	"fmt"
//	"github.com/Shopify/sarama"
//	"github.com/sirupsen/logrus"
//	"go-api/pkg/kafka"
//	"strings"
//	"sync"
//)
//
//type Kafka struct {
//	brokers           []string
//	topics            []string
//	startOffset       int64
//	version           string
//	ready             chan bool
//	group             string
//	channelBufferSize int
//	assignor          string
//}
//
//func InitKafka(cfg *kafka.Config) *Kafka {
//	return &Kafka{
//		brokers:           strings.Split(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), ","),
//		topics:            strings.Split(cfg.Topic, ","),
//		group:             cfg.Group,
//		channelBufferSize: 1000,
//		ready:             make(chan bool),
//		version:           "3.2.0",
//		assignor:          cfg.Assignor,
//	}
//}
//
//func (k *Kafka) Connect() func() {
//	logrus.Println("kafka init...")
//
//	version, err := sarama.ParseKafkaVersion(k.version)
//	if err != nil {
//		logrus.Fatalf("Error parsing Kafka version: %v", err)
//	}
//
//	config := sarama.NewConfig()
//	config.Version = version
//	// 分区分配策略
//	switch k.assignor {
//	case "sticky":
//		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
//	case "roundrobin":
//		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
//	case "range":
//		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
//	default:
//		logrus.Panicf("Unrecognized consumer group partition assignor: %s", k.assignor)
//	}
//	config.Consumer.Offsets.Initial = sarama.OffsetNewest
//	config.ChannelBufferSize = k.channelBufferSize // channel长度
//
//	// 创建client
//	newClient, err := sarama.NewClient(k.brokers, config)
//	if err != nil {
//		logrus.Fatal(err)
//	}
//	// 获取所有的topic
//	topics, err := newClient.Topics()
//	if err != nil {
//		logrus.Fatal(err)
//	}
//	logrus.Infof("topics: %#v", topics)
//
//	// 根据client创建consumerGroup
//	client, err := sarama.NewConsumerGroupFromClient(k.group, newClient)
//	if err != nil {
//		logrus.Fatalf("Error creating consumer group client: %v", err)
//	}
//
//	ctx, cancel := context.WithCancel(context.Background())
//	wg := &sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		for {
//			if err := client.Consume(ctx, k.topics, k); err != nil {
//				// 当setup失败的时候，error会返回到这里
//				logrus.Errorf("Error from consumer: %v", err)
//				return
//			}
//			// check if context was cancelled, signaling that the consumer should stop
//			if ctx.Err() != nil {
//				logrus.Println(ctx.Err())
//				return
//			}
//			k.ready = make(chan bool)
//		}
//	}()
//	<-k.ready
//	logrus.Infoln("Sarama consumer up and running!...")
//	// 保证在系统退出时，通道里面的消息被消费
//	return func() {
//		logrus.Info("kafka close")
//		cancel()
//		wg.Wait()
//		if err = client.Close(); err != nil {
//			logrus.Errorf("Error closing client: %v", err)
//		}
//	}
//}
//
//// Setup is run at the beginning of a new session, before ConsumeClaim
//func (k *Kafka) Setup(session sarama.ConsumerGroupSession) error {
//	logrus.Info("setup")
//	session.ResetOffset("mytopic", 0, 0, "")
//	// Mark the consumer as ready
//	close(k.ready)
//	return nil
//}
//
//// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
//func (k *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
//	logrus.Info("cleanup")
//	return nil
//}
//
//// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
//func (k *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//	logrus.Infof(
//		"开始消费,topic : %#v, Partition:%v, offset:%d, data: %#v",
//		claim.Topic(),
//		claim.Partition(),
//		claim.InitialOffset(),
//		len(claim.Messages()),
//	)
//	// NOTE:
//	// Do not move the code below to a goroutine.
//	// The `ConsumeClaim` itself is called within a goroutine, see:
//	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
//	// 具体消费消息
//	for message := range claim.Messages() {
//		logrus.Infof(
//			"[topic:%s] [partiton:%d] [offset:%d] [value:%s] [time:%v]",
//			message.Topic,
//			message.Partition,
//			message.Offset,
//			string(message.Value),
//			message.Timestamp,
//		)
//		// 更新位移
//		session.MarkMessage(message, "")
//	}
//	return nil
//}
