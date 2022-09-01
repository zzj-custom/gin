package mysql

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
	"sync"
	"time"
)

var (
	dbOnce  sync.Once
	clients map[string]*gorm.DB
	err     error
)

func GetConnects() map[string]*gorm.DB {
	return clients
}

func GetConnect(key string) (*gorm.DB, error) {
	if clients == nil {
		return nil, errors.New("数据库连接未初始化")
	}
	con, ok := clients[key]
	if !ok {
		return nil, errors.New("未定义的数据库连接")
	}
	return con, nil
}

func NewClient(dbs map[string]*Database) (cons map[string]*gorm.DB, err error) {
	cons = map[string]*gorm.DB{}
	for k, v := range dbs {
		var dsn = build(v)
		dialector := mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         255,
			SkipInitializeWithVersion: false,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			DontSupportForShareClause: true,
		})
		cfg := &gorm.Config{
			DisableAutomaticPing: false,
		}
		if v.UseLog {
			var logLv logger.LogLevel
			if v.LogLevel > 0 && v.LogLevel <= 4 {
				logLv = logger.LogLevel(v.LogLevel)
			} else {
				logLv = logger.Error
			}
			cfg.Logger = logger.Default.LogMode(logLv)
		}
		con, err := gorm.Open(dialector, cfg)
		if err != nil {
			logrus.Error("数据库链接创建失败", k)
			continue
		}
		sqlDb, err := con.DB()
		if err != nil {
			logrus.Error("数据库服务创建失败", k)
			continue
		}
		sqlDb.SetMaxIdleConns(v.MaxIdleConn)
		sqlDb.SetMaxOpenConns(v.MaxOpenConn)
		sqlDb.SetConnMaxIdleTime(time.Duration(v.ConnMaxFreeTime) * time.Second)
		sqlDb.SetConnMaxLifetime(time.Duration(v.ConnMaxLifeTime) * time.Second)

		cons[k] = con
	}
	return
}

func Client(dbs map[string]*Database) (map[string]*gorm.DB, error) {
	dbOnce.Do(func() {
		clients, err = NewClient(dbs)
	})
	return clients, err
}

func build(db *Database) string {
	var dsn string
	if db.DSN != "" {
		dsn = db.DSN
	} else {
		// "root:@tcp(127.0.0.1:3306)/go?parseTime=true&loc=Local"
		dsn = fmt.Sprintf(
			"%s:%s:@tcp(%s:%d)/%s?parseTime=true&loc=Local",
			url.QueryEscape(db.Username),
			url.QueryEscape(db.Password),
			db.Host,
			db.Port,
			db.Database,
		)
	}
	return dsn
}
