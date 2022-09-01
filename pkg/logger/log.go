package logger

import (
	"bytes"
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/huandu/go-tls"
	"github.com/sirupsen/logrus"
	//"go-api/config"
	"os"
	"path/filepath"
	"strings"
)

type MyFormatter struct {
}

func (receiver *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	dump.P(entry.Data)
	timeStamp := entry.Time.Format("2006-01-02 15:04:05")
	log := fmt.Sprintf(
		"%s [%s/%s:%d][GOID:%d][%s] %s\n",
		timeStamp,
		filepath.Base(entry.Caller.File),
		entry.Caller.Function,
		entry.Caller.Line,
		tls.ID(),
		strings.ToUpper(entry.Level.String()),
		entry.Message,
	)
	b.WriteString(log)
	return b.Bytes(), nil
}

// InitLogger 初始化logger
func InitLogger() *logrus.Logger {
	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = os.Stdout

	//设置日志级别
	logger.SetLevel(logrus.InfoLevel)

	// 设置Caller(如果这里不设置，那么后面使用entry.Caller获取到就是nil)
	logger.SetReportCaller(true)

	//设置日志格式
	logger.SetFormatter(&MyFormatter{})

	return logger
}

// SubscribeLog 订阅 警告日志
//func SubscribeLog(entry *logrus.Entry, subMap SubscribeMap) {
//	logger := entry.Logger
//	logger.AddHook(newSubScribeHook(subMap))
//	fmt.Println("日志订阅成功")
//}
