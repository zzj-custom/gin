package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {
	// 通过门面模式，隐藏下单过程中，后端多个系统的复杂交互
	facade := NewFacade()
	fmt.Println(facade.CreateOrder("张三", "笔记本电脑", 1))
}
