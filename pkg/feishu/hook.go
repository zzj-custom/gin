package feishu

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func (fh *FsHook) Levels() []logrus.Level {
	return fh.Level
}

type Response struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}

func (fh *FsHook) Fire(e *logrus.Entry) error {
	msg, err := getFsMessage(e)
	if err != nil {
		return err
	}

	if len(fh.Url) == 0 || len(fh.WebHookUuid) == 0 || len(fh.Secret) == 0 {
		fmt.Printf("飞书告警未配置！！！\n")
		return errors.New("飞书告警未配置")
	}

	// 配置数据
	t := time.Now()
	// 签名配置
	sign, err := getSign(t, fh.Secret)
	if err != nil {
		fmt.Printf("获取sign失败\n")
		return err
	}
	// 数据配置
	fsMsg := FsMsg{
		MsgType:   "text",
		Timestamp: strconv.FormatInt(t.Unix(), 10),
		Sign:      sign,
	}
	fsMsg.Content.Text = fmt.Sprintf(
		"name:%s\n%s",
		fh.AppName,
		msg,
	)
	requestParams, err := json.Marshal(fsMsg)
	if err != nil {
		fmt.Printf("marshal失败, err : %v\n", err)
	}

	// 数据请求
	response, err := http.Post(
		fmt.Sprintf("%s/%s", fh.Url, fh.WebHookUuid),
		"application/json",
		strings.NewReader(string(requestParams)),
	)
	if err != nil {
		fmt.Printf("数据请求失败，err:%v\n", err)
		return err
	}
	if response.StatusCode != http.StatusOK {
		fmt.Printf("请求code不为200， code: %d", response.StatusCode)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("关闭失败")
		}
	}(response.Body)
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("读取body失败，err: %v\n", res)
		return err
	}
	var r Response
	err = json.Unmarshal(res, &r)
	if err != nil {
		fmt.Printf("解析数据失败，err: %v\n", res)
		return err
	}
	if r.StatusCode != 0 || r.StatusMessage != "success" {
		fmt.Println(fmt.Sprintf("飞书告警异常, code: %d, msg: %s", r.StatusCode, r.StatusMessage))
		return errors.New(fmt.Sprintf("飞书告警异常, code: %d, msg: %s", r.StatusCode, r.StatusMessage))
	}
	return nil
}

func getSign(tm time.Time, secret string) (string, error) {
	strToSign := fmt.Sprintf(
		"%d\n%s",
		tm.Unix(),
		secret,
	)

	var data []byte
	hStr := hmac.New(sha256.New, []byte(strToSign))
	_, err := hStr.Write(data)
	if err != nil {
		fmt.Printf("sha256加密失败， err : %v", err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(hStr.Sum(nil)), nil
}

func getFsMessage(e *logrus.Entry) (string, error) {
	var (
		message map[string]any
		msg     string
	)
	dump.P(e.String())
	msgBytes, err := e.Bytes()
	if err != nil {
		fmt.Printf("转换byte失败， err： %v\n", err)
		return "", err
	}

	err = json.Unmarshal(msgBytes, &message)
	if err != nil {
		fmt.Printf("转换json失败，err: %v\n", err)
		return "", err
	}

	for k, v := range message {
		msg += fmt.Sprintf("%s:%v\n", k, v)
	}

	dump.P(msg)
	return strings.TrimSuffix(msg, "\n"), nil
}

func (c *Config) Init() *logrus.Logger {
	//实例化
	logger := logrus.New()

	//设置输出
	logger.Out = os.Stdout

	// 获取log日志等级
	fsHook := c.newFsHook()

	// 转换level并设置输出等级
	level, err := c.logLevel()
	if err != nil {
		logger.WithField("level", c.Level).WithError(err)
		return nil
	}
	logger.SetLevel(level)

	// 设置formatter格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 数据文件名，函数名，以及行号
	logger.SetReportCaller(true)

	// 添加feishuhook
	logger.AddHook(fsHook)

	return logger
}

func (c Config) logLevel() (logrus.Level, error) {
	return logrus.ParseLevel(c.Level)
}

func (c *Config) getFsHookLevel() []logrus.Level {
	var logLevel []logrus.Level

	// 转换level
	level, err := c.logLevel()
	if err != nil {
		logrus.WithField("level", c.Level).WithError(err)
		return nil
	}

	for _, allLevel := range logrus.AllLevels {
		if allLevel <= level {
			logLevel = append(logLevel, allLevel)
		}
	}
	return logLevel
}

var (
	fsHook     *FsHook
	fsHookOnce sync.Once
)

func (c *Config) newFsHook() *FsHook {
	fsHookOnce.Do(func() {
		level := c.getFsHookLevel()

		fsHook = &FsHook{
			Url:         c.Url,
			WebHookUuid: c.WebHookUuid,
			Secret:      c.Secret,
			Level:       level,
			AppName:     c.AppName,
		}
	})
	return fsHook
}
