package logger

//import (
//	"github.com/rifflock/lfshook"
//	"github.com/sirupsen/logrus"
//	//"go-api/config"
//	"time"
//)
//
//func initFileLoggerHook(logPath, module string) logrus.Hook {
//	writer := New(logPath, module, time.Hour*2)
//	lfsHook := lfshook.NewHook(lfshook.WriterMap{
//		logrus.DebugLevel: writer,
//		logrus.InfoLevel:  writer,
//		logrus.WarnLevel:  writer,
//		logrus.ErrorLevel: writer,
//		logrus.FatalLevel: writer,
//		logrus.PanicLevel: writer,
//	}, &MyFormatter{})
//
//	// writer 生成新的log文件类型 writer  在通过new hook函数 消费 fire 函数
//	// writer 是实现了writer 接口的库，在日志调用write是做预处理
//	return lfsHook
//}
//
//// FileLoggerHook 文件日志
//func FileLoggerHook(module, appField string) *logrus.Entry {
//	logger := InitLogger()
//	//自定writer就行， hook 交给 lfshook
//	logger.AddHook(initFileLoggerHook(config.Config().Logger.Filepath, module))
//	return logger.WithFields(logrus.Fields{
//		"app": appField,
//	})
//}
