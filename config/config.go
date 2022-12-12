package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"go-api/pkg/email"
	"go-api/pkg/feishu"
	"go-api/pkg/jwt"
	"go-api/pkg/kafka"
	"go-api/pkg/logger"
	"go-api/pkg/mysql"
	"go-api/pkg/nacos"
	"go-api/pkg/redis"
	"path/filepath"
	"sync"
	"time"
)

var (
	globalConfig *config
	configOnce   sync.Once
	Toml         = "./config.toml"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type serverConfig struct {
	Host              string   `toml:"host"`
	ReadTimeout       duration `toml:"read_timeout"`
	ReadHeaderTimeout duration `toml:"read_header_timeout"`
	WriteTimeout      duration `toml:"write_timeout"`
	AppName           string   `toml:"app_name"`
}

func FsLog() *logrus.Logger {
	fsConfig := globalConfig.Feishu
	fsConfig.AppName = globalConfig.Server.AppName
	return fsConfig.Init()
}

type config struct {
	Server   serverConfig               `toml:"server"`
	Database map[string]*mysql.Database `toml:"database"`
	Redis    *redis.DialConfig          `toml:"redis"`
	Jwt      *jwt.Config                `toml:"jwt"`
	Email    *email.Config              `toml:"email"`
	Logger   *logger.FileConfig         `toml:"logger"`
	Kafka    *kafka.Config              `toml:"kafka"`
	Feishu   *feishu.Config             `toml:"feishu"`
	Nacos    *nacos.Config              `toml:"nacos"`
}

func ParseConfigFile(cfgFile string) *config {
	if cfgFile != "" {
		Toml = cfgFile
	}
	configOnce.Do(func() {
		globalConfig = loadConfig(Toml)
	})
	return globalConfig
}

func loadConfig(configFile string) *config {
	filePath, err := filepath.Abs(configFile)
	if err != nil {
		logrus.Panicf("解析配置文件出现异常，配置文件: %s，异常：%s\n", configFile, err)
	}
	cfg := new(config)
	if _, err := toml.DecodeFile(configFile, cfg); err != nil {
		logrus.Panicf("解析文件失败， 错误：%v\n", err)
	}
	logrus.Infof("解析配置文件完成：%s\n", filePath)
	return cfg
}

func Config() *config {
	return globalConfig
}
