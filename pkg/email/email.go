// Package email 邮件发送，使用选项设计模式
package email

import (
	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
	"gopkg.in/errgo.v2/fmt/errors"
	"net/textproto"
	"regexp"
	"time"
)

func (c *Config) initOptions(to []string, subject string, opts ...Option) {
	options := &Options{}
	for _, opt := range opts {
		opt(options)
	}
	// 邮箱检查
	mailTo := make([]Receiver, 0, len(to))
	for _, mt := range to {
		mailTo = append(mailTo, Receiver{Email: mt})
	}
	options = options.ReceiversCheck(mailTo)
	if nil == options {
		return
	}
	options.subject = subject
	c.options = options
}

func WithTextOptions(text string) Option {
	return func(opt *Options) {
		opt.text = []byte(text)
	}
}

func WithHtmlOptions(html string) Option {
	return func(opt *Options) {
		opt.html = []byte(html)
	}
}

func WithCarbonCopyOptions(cc []string) Option {
	return func(opt *Options) {
		opt.carbonCopy = cc
	}
}

func WithAttachFile(af []string) Option {
	return func(opt *Options) {
		opt.attachFile = af
	}
}

func (r *Receiver) Check() bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(r.Email)
}

func (op *Options) ReceiverCheck(email string) *Options {
	rev := &Receiver{Email: email}
	if rev.Check() {
		op.mailTo = []string{email}
		return op
	}
	logrus.WithField("email", email).WithError(
		errors.Newf("【邮箱：%s】格式有误，请重新输入", email),
	)
	return nil
}

func (op *Options) ReceiversCheck(receivers []Receiver) *Options {
	for _, rec := range receivers {
		if rec.Check() {
			op.mailTo = append(op.mailTo, rec.Email)
		} else {
			logrus.WithField("email", rec.Email).WithError(
				errors.Newf("email check fail 【%s】\n", rec.Email),
			)
			return nil
		}
	}
	return op
}

func (c *Config) initEmail() (*email.Email, error) {
	// 获取配置信息
	options := c.options

	e := &email.Email{
		Headers: textproto.MIMEHeader{},
		// 设置发送方的邮箱
		From: c.From,
		// 设置接收方的邮箱
		To: options.mailTo,
		// 设置主题
		Subject: options.subject,
		// 设置发送的内容
		Text: options.text,
		// 发送html
		HTML: options.html,
		// 抄送
		Cc: options.carbonCopy,
	}

	// 批量发送附件
	for _, o := range options.attachFile {
		_, err := e.AttachFile(o)
		if err != nil {
			logrus.Panicf("添加附件失败，错误：%v", err)
			return nil, err
		}
	}

	return e, nil
}

func (c *Config) SendMail(to []string, subject string, opts ...Option) error {
	// 初始化请求参数
	c.initOptions(to, subject, opts...)

	// 初始化email
	e, err := c.initEmail()
	if err != nil {
		return err
	}

	// 配置服务器信息
	err = pool().Send(e, 10*time.Second)
	if err != nil {
		logrus.Panicf("发送邮箱失败, 错误：%v", err)
		return err
	}
	logrus.Info("邮箱发送成功")
	return nil
}
