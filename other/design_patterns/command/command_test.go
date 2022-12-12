package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	//创建电饭煲命令接受者
	electricCooker := new(ElectricCooker)
	// 创建电饭煲指令接收器
	electricCookerInvoker := new(ElectricCookerInvoker)

	// 蒸饭
	streamRiceCommand := NewSteamRiceCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(streamRiceCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 煮粥
	cookCongeeCommand := NewCookCongeeCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(cookCongeeCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())

	// 停止命令
	shutdownCommand := NewShutdownCommand(electricCooker)
	electricCookerInvoker.SetCookCommand(shutdownCommand)
	fmt.Println(electricCookerInvoker.ExecuteCookCommand())
}
