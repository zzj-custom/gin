package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"go-api/app/middleware"
	cmdConfig "go-api/cmd/config"
	"go-api/cmd/router"
	"go-api/cmd/validate"
	"go-api/config"
	"go-api/event/service"
	"go-api/internal/task"
	"go-api/pkg/email"
	"go-api/pkg/mysql"
	"go-api/pkg/nacos"
	"go-api/pkg/redis"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func usage() {
	fmt.Printf("Usage: %s [-c CONFIG_FILE]\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&config.Toml, "c", "../config/config.toml", "配置文件路径")
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 解析配置文件
	cfg := config.ParseConfigFile(config.Toml)

	// 初始化邮件依赖项
	emailPool := email.InitPool(cfg.Email)

	// 注册nacos
	configHandler := new(cmdConfig.Handler)
	nacos.RegisterHandler(configHandler)
	//config.Config().Nacos.GetAllConfigFile()

	// 初始化数据库依赖项
	dbs, err := mysql.Client(cfg.Database)
	if err != nil {
		logrus.WithError(err).Error("初始化数据库连接失败，请检查网络连接和数据库配置")
		return
	}
	// 获取所有的model
	//for k, db := range dbs {
	//	config.AutoMigrate(db, cfg.Database[k].ModelPath)
	//}

	// 初始化redis依赖项
	redisPool := redis.InitPool(cfg.Redis)

	// 关闭链接
	defer func() {
		logrus.Info("关闭数据库链接")
		for _, db := range dbs {
			if conn, err := db.DB(); err != nil {
				_ = conn.Close()
			}
		}

		logrus.Info("关闭redis链接")
		_ = redisPool.Close()

		logrus.Info("关闭邮箱链接")
		emailPool.Close()

		logrus.Info("服务已经关闭")

	}()

	// 注册新的V10版本校验器
	binding.Validator = new(validate.DefaultValidator)

	// 初始化gin
	gin.SetMode(gin.ReleaseMode)
	e := gin.Default()
	e.Use(middleware.Log)
	for _, r := range router.Routers() {
		r(e)
	}

	// 启动监听
	srv := &http.Server{
		Addr:              cfg.Server.Host,
		Handler:           e,
		ReadTimeout:       cfg.Server.ReadTimeout.Duration,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout.Duration,
		WriteTimeout:      cfg.Server.WriteTimeout.Duration,
	}
	srv.SetKeepAlivesEnabled(true)

	// 开启事件监听
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	// 初始化taskPool
	taskPool := task.Pool()
	defer taskPool.Release()

	// 开启事件监听
	go service.Start(cancelCtx)

	// 启动协程运行服务监听
	logrus.Infof("启动服务，监听地址：%s\n", cfg.Server.Host)
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {

			logrus.Fatalf("启动服务失败或服务运行异常， 错误：%s", err)
		}
	}()

	// 启动信号监听实现任务阻塞
	// 信号：0x02(INT), 0x09(KILL), 0x0F(TERM)
	quite := make(chan os.Signal)
	signal.Notify(quite, os.Interrupt, os.Kill, os.Signal(syscall.SIGTERM))
	sig := <-quite

	// 收到监听的信号，准备退出程序
	logrus.Infof("收到退出的信号为：%d", sig)
	cancelFunc()

	// 加入上下文超时处理
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("关闭服务失败，错误：%v", err)
	}
	logrus.Info("Web服务关闭，开始回收数据库连接等资源")
}
