// Package design_patterns 设计模式

// 1. 命令模式是一种行为设计模式，它可将请求转换为一个包含与请求相关的所有信息的独立对象。
// 该转换让你能根据不同的请求将方法参数化、延迟请求执行或将其放入队列中，且能实现可撤销操作。

// 2. 方法参数化是指将每个请求参数传入具体命令的工厂方法（go语言没有构造函数）创建命令，同时具体命令会默认设置好接受对象，
// 这样做的好处是不管请求参数个数及类型，还是接受对象有几个，都会被封装到具体命令对象的成员字段上，并通过统一的Execute接口方法进行调用，屏蔽各个请求的差异，便于命令扩展，多命令组装，回滚等；

package command

// example 控制电饭煲做饭是一个典型的命令模式的场景，电饭煲的控制面板会提供设置煮粥、蒸饭模式，及开始和停止按钮，
// 电饭煲控制系统会根据模式的不同设置相应的火力，压强及时间等参数；煮粥，蒸饭就相当于不同的命令，开始按钮就相当命令触发器，设置好做饭模式，点击开始按钮电饭煲就开始运行，同时还支持停止命令；

import "fmt"

// ElectricCooker 电饭煲
type ElectricCooker struct {
	fire     string // 火力
	pressure string // 压力
}

// SetFire 设置火力
func (e *ElectricCooker) SetFire(fire string) {
	e.fire = fire
}

// SetPressure 设置压力
func (e *ElectricCooker) SetPressure(pressure string) {
	e.pressure = pressure
}

// Run 持续运行指定时间
func (e *ElectricCooker) Run(duration string) string {
	return fmt.Sprintf("电饭煲设置火力为%s,压力为%s,持续运行%s;", e.fire, e.pressure, duration)
}

// Shutdown 停止
func (e *ElectricCooker) Shutdown() string {
	return "电饭煲停止运行。"
}
