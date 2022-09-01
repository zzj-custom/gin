package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/jsonutil"
	"github.com/sirupsen/logrus"
	"go-api/internal/util"
	"go-api/pkg/kafka"
	"strconv"
	"strings"
)

func Producer(cfg *kafka.Config, value map[string]string) (int32, int64) {
	// 获取基础配置
	config := kafka.GetConfig(cfg)

	// 构建消息
	v, _ := jsonutil.Encode(value)
	msg := &sarama.ProducerMessage{
		Topic: cfg.Topic,
		Key:   sarama.StringEncoder(strconv.FormatInt(util.Get(1), 10)),
		Value: sarama.StringEncoder(v),
	}

	// 连接kafka
	host := strings.Split(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), ",")
	client := kafka.ProducerConnect(host, config)
	defer client.Close()

	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	dump.P(err)
	if err != nil {
		logrus.Infof("发送消息失败, 错误：%v", err)
		return 0, 0
	}
	return pid, offset
}
