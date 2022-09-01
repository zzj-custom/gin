package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-api/internal/consts"
	"go-api/internal/util"
	"net/http"
)

// ReformatBody 格式化内容体中间件
// 当且仅当context.Set("__CONTEXT_OUTPUT_RESULT__", interface{})时才正常地格式化
// 否则会输出200响应码的json错误应答
func ReformatBody(ctx *gin.Context) {
	ctx.Next()
	if ctx.Writer.Written() {
		log.WithField("request", ctx.Request).Error("http请求头已经发出")
		return
	}
	r, ok := ctx.Get(consts.ContextResult)
	if !ok {
		ctx.JSON(http.StatusOK, util.Error(consts.InvalidResponse))
		return
	}
	if ri, ok := r.(int); ok {
		ctx.JSON(http.StatusOK, util.Error(ri))
		return
	}
	ctx.JSON(http.StatusOK, util.Ok(r))
}
