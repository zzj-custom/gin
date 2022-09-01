package login

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-api/config"
	"go-api/event/dispatcher"
	"go-api/event/service/email_event"
	"go-api/internal/consts"
	"go-api/internal/model"
	"go-api/internal/model/login"
	"go-api/internal/util"
	"go-api/pkg/jwt"
	"strconv"
)

type RegisterRequestDTO struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}

type RegisterResponseDTO struct {
	Result string `json:"result"`
}

type UserLoginRequestDTO struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type UserLoginResponseDTO struct {
	Token     string      `json:"token"`
	UserName  string      `json:"userName"`
	SecretKey string      `json:"secretKey"`
	AppId     int64       `json:"appId"`
	Email     string      `json:"email"`
	Phone     int64       `json:"phone"`
	CreateAt  model.XTime `json:"createAt"`
}

func Register(req RegisterRequestDTO) (*RegisterResponseDTO, error) {
	password := jwt.SecretSalt(req.Password, config.Config().Jwt.Salt)
	// 判断账号是否存在
	userRepo := login.NewUserRepo()
	userFindRes := userRepo.FindUserInfoByUserName(
		req.UserName,
		"",
	)
	if userFindRes.ID != 0 {
		logrus.Infof("账户已经存在，data:%v", userFindRes)
		return nil, util.ErrorRes(consts.AccountIsExists)
	}

	// 创建数据
	_, err := userRepo.CreateUser(req.UserName, password, req.Ip)
	if err != nil {
		logrus.Infof("账号创建失败，错误：%v", err)
		return nil, util.ErrorRes(consts.AccountCreateFail)
	}
	resStr := fmt.Sprintf("恭喜您：%s，注册成功", req.UserName)

	// 发送邮件
	f := dispatcher.NewManagerEvent().RegisterFeed(email_event.Email)
	evt := &email_event.EmailInfo{
		MailTo:  []string{"1844066417@qq.com"},
		Subject: "注册账号",
		Text:    resStr,
	}
	f.Send(evt)

	return &RegisterResponseDTO{
		Result: resStr,
	}, nil
}

func UserLogin(req UserLoginRequestDTO) (*UserLoginResponseDTO, error) {
	// 验证账号密码
	userRepo := login.NewUserRepo()
	password := jwt.SecretSalt(req.Password, config.Config().Jwt.Salt)
	userInfo := userRepo.FindUserInfoByUserName(req.UserName, password)
	if userInfo.ID == 0 {
		return nil, util.ErrorRes(consts.AccountNotFound)
	}
	// 返回给前端token
	//userId, _ := util.AesCtrCrypt([]byte(strconv.FormatInt(int64(userInfo.ID), 10)))
	tokenEncode := jwt.Token{
		UserName:  userInfo.UserName,
		SecretKey: userInfo.SecretKey,
		AppId:     strconv.FormatInt(userInfo.AppId, 10),
		Email:     userInfo.Email,
		Phone:     strconv.FormatInt(userInfo.Phone, 10),
	}

	token, err := jwt.GetToken(userInfo.ID, tokenEncode, []byte(config.Config().Jwt.TokenSecret))
	if err != nil {
		logrus.Infof("token生成失败，错误：%v", err)
		return nil, util.ErrorRes(consts.AccountCreateFail)
	}
	return &UserLoginResponseDTO{
		Token:     token,
		UserName:  userInfo.UserName,
		SecretKey: userInfo.SecretKey,
		AppId:     userInfo.AppId,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		CreateAt:  userInfo.CreatedAt,
	}, nil
}
