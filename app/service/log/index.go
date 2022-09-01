package log

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//	"github.com/sirupsen/logrus/hooks/test"
//	"go-api/internal/model/log"
//	"time"
//)
//
//func CreateLog(ctx *gin.Context, res string) {
//	logger, hook := test.NewNullLogger()
//	s := hook.LastEntry().Level.String()
//	_, err := logrus.LastEntry()
//	if err != nil {
//		logrus.Info("获取日志级别失败")
//	}
//	err = handleCreate(
//		ctx.GetString("request_id"),
//		ctx.Request.Proto,
//		ctx.Request.URL.Path,
//		res,
//		200,
//	)
//	if err != nil {
//		logrus.Info("创建日志失败")
//	}
//}
//
//func handleCreate(requestId string, channel string, message string, context string, level int) error {
//	repo := log.NewLogsRepo()
//	res, err := repo.CreateLog(
//		requestId,
//		channel,
//		message,
//		context,
//		level,
//		time.Now(),
//	)
//	if err != nil && res.ID == 0 {
//		logrus.Infof(
//			"日志创建失败，【request_id:%s, channel:%s, message:%s】 错误：%v",
//			requestId,
//			channel,
//			message,
//			err,
//		)
//		return err
//	}
//	return nil
//}
