package test

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"go-api/config"
	"go-api/pkg/email"
	"testing"
)

func TestSendMail(t *testing.T) {
	resStr := fmt.Sprintf("恭喜您：%s，注册成功", "邹祝家")
	// 初始化参数
	el := config.ParseConfigFile("./config.toml").Email
	email.InitPool(el)
	err := el.SendMail(
		[]string{"1844066417@qq.com"},
		"注册账号",
		email.WithTextOptions(resStr),
		email.WithHtmlOptions(resStr),
		email.WithCarbonCopyOptions([]string{}),
		email.WithAttachFile([]string{}),
	)
	assert.Equal(t, err, nil)
}

func TestStruct(t *testing.T) {
	//managerEvent := dispatcher.NewManagerEvent()
	//managerEvent.SetListens(email_event.EmailEvent{})
	//listens := managerEvent.GetListens()
	//for _, listen := range listens {
	//	voh := reflect.TypeOf(listen)
	//	if voh.Kind() != reflect.Struct {
	//		fmt.Printf("数据：%v, 不是结构体", listen)
	//	}
	//	methodByName, ok := voh.MethodByName("Run")
	//	if !ok {
	//		fmt.Printf("结构体：%v, 没有Run方法", listen)
	//		continue
	//	}
	//	methodByName.Func.Call([]reflect.Value{reflect.ValueOf(listen), reflect.ValueOf(context.Background())})
	//}
}
