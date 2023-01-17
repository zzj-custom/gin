package util

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"testing"
)

func TestAdd(t *testing.T) {
	a := []int{1, 2, 5, 6}
	b := []int{3, 4}
	//dump.P(Add(a, b, 2))
	copy(a[2:], b)
	dump.P(a)
}

func TestDelete(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	dump.P(Delete(a, 2, 2))
}

// 可以使用无缓存通道来实现协程同步
func TestChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		fmt.Println("hello world")
		fmt.Println("开始")
		ch <- 1
	}()
	// 阻塞
	<-ch
	fmt.Println("结束")
}
