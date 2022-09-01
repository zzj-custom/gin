package util

import (
	"github.com/gin-gonic/gin"
)

func GetIP(ctx *gin.Context) string {
	forwarded := ctx.Request.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return ctx.Request.RemoteAddr
}
