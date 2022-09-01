package log

import (
	"github.com/sirupsen/logrus"
	"go-api/internal/model"
	"go-api/pkg/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Logs struct {
	model.Model
	RequestId string      `gorm:"column:request_id;type:varchar(64);comment:请求ID;not null;index"`
	Channel   string      `gorm:"column:channel;comment:渠道;type:varchar(64);"`
	message   string      `gorm:"column:message;comment:消息;type:varchar(128)"`
	context   string      `gorm:"column:context;comment:内容;type:longtext"`
	level     int         `gorm:"column:level;comment:错误级别;type:int"`
	logAt     model.XTime `gorm:"column:log_at;comment:日志时间;type:timestamp"`
}

var (
	logRepo     *logsRepository
	logRepoOnce sync.Once
)

type logsRepository struct {
	db *gorm.DB
}

func (l Logs) name() string {
	return "log"
}

func NewLogsRepo() *logsRepository {
	logRepoOnce.Do(func() {
		logRepo = new(logsRepository)
		db, err := mysql.GetConnect("go")
		if err != nil {
			logrus.Fatal("log连接失败")
			return
		}
		logRepo.db = db
	})
	return logRepo
}

func (r *logsRepository) CreateLog(requestId string, channel string, message string, context string, level int, logAt time.Time) (Logs, error) {
	var record = Logs{
		RequestId: requestId,
		Channel:   channel,
		message:   message,
		context:   context,
		level:     level,
		logAt:     model.XTime{Time: logAt},
	}
	tx := r.db.Create(&record)
	return record, tx.Error
}
