package login

import (
	"github.com/gin-gonic/gin"
	"go-api/app/requests"
	"go-api/app/service/login"
	"go-api/internal/consts"
	"go-api/internal/util"
)

func Register(ctx *gin.Context) {
	var req requests.RegisterRequest
	if err := ctx.ShouldBind(&req); err != nil {
		util.ResponseWrapper(ctx, util.ErrorRes(consts.InvalidData), nil)
		return
	}
	reqDto := login.RegisterRequestDTO{
		UserName: req.UserName,
		Password: req.Password,
		Ip:       util.GetIP(ctx),
	}
	resDto, err := login.Register(reqDto)
	util.ResponseWrapper(ctx, err, resDto)
}

func UserLogin(ctx *gin.Context) {
	var req requests.UserLoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		util.ResponseWrapper(ctx, util.ErrorRes(consts.InvalidData), nil)
		return
	}
	reqDto := login.UserLoginRequestDTO{
		UserName: req.UserName,
		Password: req.Password,
	}
	resDto, err := login.UserLogin(reqDto)
	util.ResponseWrapper(ctx, err, resDto)
}
