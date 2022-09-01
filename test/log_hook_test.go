package test

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestGetNewFieldLoggerContext(t *testing.T) {
	//lg := logger.FileLoggerHook("test", "go-api")
	//lg.Info("info消息日志1")
}

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Unix())
}

func TestFsHook(t *testing.T) {
	//feishu.Init().Info("测试数据")
	logrus.Info("测试")
}
