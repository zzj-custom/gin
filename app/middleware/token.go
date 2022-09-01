package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/jsonutil"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"go-api/config"
	"go-api/internal/consts"
	"go-api/internal/util"
	"go-api/pkg/jwt"
	"net/http"
	"strings"
)

func Token(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		logrus.Info("token 不存在")
		ctx.JSON(http.StatusUnauthorized, util.Error(consts.InvalidToken))
		ctx.Abort()
		return
	}

	// 按空格分隔
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		logrus.Infof("【token：%s】格式有问题", token)
		ctx.JSON(http.StatusUnauthorized, util.Error(consts.InvalidToken))
		ctx.Abort()
		return
	}

	tok, err := jwt.ParseToken(token, []byte(config.Config().Jwt.TokenSecret))
	if err != nil {
		logrus.Infof("【token：%v】无效，错误：%s", tok, err)
		ctx.JSON(http.StatusUnauthorized, util.Error(consts.InvalidToken))
		ctx.Abort()
		return
	}
	userInfo := make(map[string]string)
	err = jsonutil.Decode([]byte(tok.UserName), &userInfo)
	if err != nil {
		logrus.Infof("【token：%v】无法解析，错误：%s", tok, err)
		ctx.JSON(http.StatusUnauthorized, util.Error(consts.InvalidToken))
		ctx.Abort()
		return
	}
	// 将当前请求的username信息保存到请求的上下文c上
	ctx.Set("user_info", userInfo)
	ctx.Set("user_id", tok.Id)
	ctx.Set("request_id", uuid.NewV4())
	ctx.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
}
