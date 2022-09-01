package data_structure

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"sync"
	"time"
)

type MessageQueue interface {
	Send(message int)
	Handler()
	Size() int
	Capacity() int
	Get() chan int
}

type Message struct {
	queue    chan int
	capacity int
	once     sync.Once
	closed   bool
	sync.Mutex
}

func (mc *Message) Get() chan int {
	return mc.queue
}

func (mc *Message) Send(message int) {
	if mc.IsClosed() {
		fmt.Printf("通道已经被关闭\n")
		return
	}
	mc.Lock()
	select {
	case mc.queue <- message:
		fmt.Printf("数据存入，msg: %v\n", message)
		if len(mc.queue) == mc.capacity {
			fmt.Printf("数据达到上线，len: %d, cap:%d\n", len(mc.queue), mc.capacity)
			mc.once.Do(func() {
				if !mc.closed {
					close(mc.queue)
					mc.closed = true
				}
			})
		}
	default:

	}
	mc.Unlock()
}

func (mc *Message) IsClosed() bool {
	return mc.closed
}

func (mc *Message) Handler() {
	for {
		select {
		case msg, ok := <-mc.queue:
			dump.P(msg, ok)
			if ok {
				fmt.Printf("开始处理数据, msg:%v\n", msg)
			} else {
				fmt.Printf("通道关闭\n")
				return
			}
		case <-time.After(10 * time.Second):
			fmt.Printf("处理超时")
		}
	}
}

func (mc *Message) Size() int {
	return len(mc.queue)
}

func (mc *Message) Capacity() int {
	return mc.capacity
}

func NewMessage(capacity int) MessageQueue {
	var mc MessageQueue
	mc = &Message{
		queue:    make(chan int, capacity),
		capacity: capacity,
	}
	return mc
}
