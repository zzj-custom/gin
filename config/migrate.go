package config

import (
	"github.com/sirupsen/logrus"
	"go-api/internal/model/log"
	"go-api/internal/model/login"
	"go-api/internal/util"
	"gorm.io/gorm"
)

func initMigrate() map[string]func() interface{} {
	return map[string]func() interface{}{
		"user": func() interface{} {
			return login.User{}
		},
		"logs": func() interface{} {
			return log.Logs{}
		},
	}
}

func handleMigrate(db *gorm.DB, table string) {
	AutoMigrateFunc := initMigrate()
	if db.Migrator().HasTable(table) {
		logrus.Infof("数据表【%s】已存在\n", table)
	}
	if value, ok := AutoMigrateFunc[table]; !ok {
		logrus.Panicf("数据表【%s】没有定义model层init方法初始化struct\n", table)
	} else {
		if err := db.AutoMigrate(value()); err != nil {
			logrus.Panicf("数据表【%s】生成失败,err: %v\n", table, err)
		}
	}
	logrus.Infof("数据表【%s】生成成功", table)

}

func AutoMigrate(db *gorm.DB, path string) {
	models, err := util.FileWalk(path)
	if err != nil {
		logrus.Panicf("路径解析失败，path:%s", path)
	}
	for _, m := range models {
		handleMigrate(db, m)
	}
}
