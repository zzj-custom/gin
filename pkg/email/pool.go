package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"sync"
)

var (
	emailPool *email.Pool
	emailOnce sync.Once
)

func pool() *email.Pool {
	if emailPool == nil {
		panic("email连接池未初始化")
	}
	return emailPool
}

func InitPool(config *Config) *email.Pool {
	emailOnce.Do(func() {
		emailPool = newPool(config)
	})
	return emailPool
}

func newPool(config *Config) *email.Pool {
	p, err := email.NewPool(
		fmt.Sprintf("%s:%d", config.Host, config.Port),
		10,
		smtp.PlainAuth("", config.Username, config.Password, config.Host),
	)
	if err != nil {
		logrus.Panicf("初始化邮箱发送失败,错误：%v", err)
		return nil
	}
	return p
}
