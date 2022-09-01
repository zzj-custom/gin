package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go-api/config"
	"go-api/internal/util"
	"time"
)

func NotifyTimeout(ctx *gin.Context) {
	// 开始时间
	start := time.Now()

	ctx.Next()

	// 结束时间
	end := time.Now()

	//计算相差时间
	latency := end.Sub(start).Microseconds()

	config.FsLog().WithFields(logrus.Fields{
		"latency":    fmt.Sprintf("%dms", latency),
		"path":       ctx.Request.URL.Path,
		"clientIP":   util.GetIP(ctx),
		"method":     ctx.Request.Method,
		"statusCode": ctx.Writer.Status(),
		"user-agent": ctx.Request.UserAgent(),
		"params":     util.GetRequestParams(ctx, ctx.Request.Method),
	}).Warn("测试")
}
