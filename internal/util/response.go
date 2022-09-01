package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/internal/consts"
	"net/http"
)

type ApiResponse interface {
	GetCode() int
	GetMessage() string
}

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r *response) GetCode() int {
	return r.Code
}

func (r response) GetMessage() string {
	return r.Msg
}

type okResponse struct {
	response
	Data interface{} `json:"data"`
}

func Ok(data interface{}) *okResponse {
	return &okResponse{
		response: response{
			Code: http.StatusOK,
			Msg:  "success",
		},
		Data: data,
	}
}

func Error(i int) *response {
	msg, ok := consts.CodeMessages[i]
	if !ok {
		msg = consts.CodeMessages[-1]
	}
	return &response{
		Code: i,
		Msg:  msg,
	}
}

func ResponseWrapper(ctx *gin.Context, err error, response interface{}) {
	if err != nil {
		var code = 99999999
		if errorResponse, ok := err.(*ErrorResponse); ok {
			code = errorResponse.Code
		}
		ctx.Set(consts.ContextResult, code)
		return
	}
	ctx.Set(consts.ContextResult, response)
}

type ErrorResponse struct {
	error
	Code int
}

func ErrorRes(code int, msg ...string) *ErrorResponse {
	var m = ""
	if len(msg) > 0 {
		m = fmt.Sprintf("%s", msg)
	}
	return &ErrorResponse{
		error: errors.New(m),
		Code:  code,
	}
}
