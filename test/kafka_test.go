package test

import (
	"github.com/gookit/goutil/dump"
	"github.com/sirupsen/logrus"
	"go-api/config"
	"go-api/pkg/kafka/producer"
	"testing"
)

func TestProducer(t *testing.T) {
	cfg := config.ParseConfigFile(config.Toml)
	if cfg == nil {
		logrus.Info("配置为空")
	}
	p, o := producer.Producer(cfg.Kafka, map[string]string{
		"id":    "12",
		"name":  "zouzhujia",
		"appId": "zqp113",
	})
	dump.Println(p, o)
}

func TestConsumer(t *testing.T) {
	cfg := config.ParseConfigFile(config.Toml)
	if cfg == nil {
		logrus.Info("配置为空")
	}
	//consumer.SaramaConsumerGroup(cfg.Kafka)
}
