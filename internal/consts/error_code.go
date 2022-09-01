package consts

import "net/http"

const (
	ContextResult     = "__CONTEXT_OUTPUT_RESULT__"
	OK                = http.StatusOK
	InvalidResponse   = http.StatusInternalServerError
	InvalidToken      = 40010001
	InvalidData       = 40010002
	AccountIsExists   = 40010003
	AccountCreateFail = 40010004
	AccountNotFound   = 40010005
	SendEmailFail     = 40010006
	SystemError       = -1
)

var CodeMessages map[int]string

func init() {
	CodeMessages = map[int]string{
		InvalidResponse:   "无效的应答数据",
		InvalidToken:      "无效的Token",
		InvalidData:       "数据错误",
		SystemError:       "system error",
		AccountIsExists:   "账户已经存在，请重新选择",
		AccountCreateFail: "账户创建失败，请联系管理员",
		AccountNotFound:   "账户不存在，请重新输入",
		SendEmailFail:     "邮件发送失败",
	}
}
