package adapter

import "fmt"

// HuaweiPlug 华为手机充电插槽接口
type HuaweiPlug interface {
	ConnectTypeC() string
}

// HuaweiPhone 华为系列手机
type HuaweiPhone struct {
	model string
}

// NewHuaweiPhone 华为手机创建方法
func NewHuaweiPhone(model string) *HuaweiPhone {
	return &HuaweiPhone{
		model: model,
	}
}

// ConnectTypeC 华为手机TypeC充电插槽
func (h *HuaweiPhone) ConnectTypeC() string {
	return fmt.Sprintf("%v connect typeC plug", h.model)
}

// ApplePlug 苹果手机充电插槽
type ApplePlug interface {
	ConnectLightning() string
}

// Apple IPhone 苹果系列手机
type Apple struct {
	model string
}

// NewApple NewIPhone 苹果手机创建方法
func NewApple(model string) *Apple {
	return &Apple{
		model: model,
	}
}

// ConnectLightning 苹果手机Lightning充电插槽
func (a *Apple) ConnectLightning() string {
	return fmt.Sprintf("%v connect lightning plug", a.model)
}
