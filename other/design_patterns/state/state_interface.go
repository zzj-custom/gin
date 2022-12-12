// Package design_patterns 设计模式-状态模式

// 状态模式是一种行为设计模式，让你能在一个对象的内部状态变化时改变其行为，使其看上去就像改变了自身所属的类一样。
// 该模式将与状态相关的行为抽取到独立的状态类中，让原对象将工作委派给这些类的实例，而不是自行进行处理。
// 状态迁移有四个元素组成，起始状态、触发迁移的事件，终止状态以及要执行的动作，每个具体的状态包含触发状态迁移的执行方法，
// 迁移方法的实现是执行持有状态对象的动作方法，同时设置状态为下一个流转状态；
// 持有状态的业务对象包含有触发状态迁移方法，这些迁移方法将请求委托给当前具体状态对象的迁移方法。

package state

import "fmt"

// example
// IPhone手机充电就是一个手机电池状态的流转，一开始手机处于有电状态，插入充电插头后，继续充电到满电状态，并进入断电保护，
// 拔出充电插头后使用手机，由满电逐渐变为没电，最终关机；

// BatteryState 电池状态接口，支持手机充电线插拔事件
type BatteryState interface {
	ConnectPlug(iPhone *IPhone) string
	DisconnectPlug(iPhone *IPhone) string
}

// fullBatteryState 满电状态
type fullBatteryState struct{}

func (s *fullBatteryState) String() string {
	return "满电状态"
}

func (s *fullBatteryState) ConnectPlug(iPhone *IPhone) string {
	return iPhone.pauseCharge()
}

func (s *fullBatteryState) DisconnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(PartBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, PartBatteryState)
}

// emptyBatteryState 空电状态
type emptyBatteryState struct{}

func (s *emptyBatteryState) String() string {
	return "没电状态"
}

func (s *emptyBatteryState) ConnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(PartBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, PartBatteryState)
}

func (s *emptyBatteryState) DisconnectPlug(iPhone *IPhone) string {
	return iPhone.shutdown()
}

// partBatteryState 部分电状态
type partBatteryState struct{}

func (s *partBatteryState) String() string {
	return "有电状态"
}

func (s *partBatteryState) ConnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(FullBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, FullBatteryState)
}

func (s *partBatteryState) DisconnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(EmptyBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, EmptyBatteryState)
}
