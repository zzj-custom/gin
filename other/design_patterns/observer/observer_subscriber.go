// Package design_patterns 设计模式-观察者模式

// 观察者模式是一种行为设计模式，允许你定义一种订阅机制，可在对象事件发生时通知多个 “观察” 该对象的其他对象。
// 观察者模式提供了一种作用于任何实现了订阅者接口的对象的机制，可对其事件进行订阅和取消订阅。
// 观察者模式是最常用的模式之一，是事件总线，分布式消息中间件等各种事件机制的原始理论基础，常用于解耦多对一的对象依赖关系；
// 增强的实现功能包括：
// 1. 当被观察者通过异步实现通知多个观察者时就相当于单进程实例的消息总线；
// 2. 同时还可以根据业务需要，将被观察者所有数据状态变更进行分类为不同的主题，观察者通过不同主题进行订阅；
// 3. 同一个主题又可分为增加，删除，修改事件行为；
// 4. 每个主题可以实现一个线程池，多个主题通过不同的线程池进行处理隔离，线程池可以设置并发线程大小、缓冲区大小及调度策略，比如先进先出，优先级等策略；
// 5. 观察者处理事件时有可能出现异常，所以也可以注册异常处理函数，异常处理也可以通过异常类型进行分类；
// 6. 根据业务需求也可以实现通知异常重试，延迟通知等功能；

package observer

import "fmt"

// example
// 信用卡业务消息提醒可通过观察者模式实现，业务消息包括日常消费，出账单，账单逾期，消息提醒包括短信、邮件及电话，
// 根据不同业务的场景会采用不同的消息提醒方式或者多种消息提醒方式，这里信用卡相当于被观察者，观察者相当于不同的通知方式；
// 日常消费通过短信通知，出账单通过邮件通知，账单逾期三种方式都会进行通知；

type Subscriber interface {
	Name() string          //订阅者名称
	Update(message string) //订阅更新方法
}

// shortMessage 信用卡消息短信订阅者
type shortMessage struct{}

func (s *shortMessage) Name() string {
	return "手机短息"
}

func (s *shortMessage) Update(message string) {
	fmt.Printf("通过【%s】发送消息:%s\n", s.Name(), message)
}

// email 信用卡消息邮箱订阅者
type email struct{}

func (e *email) Name() string {
	return "电子邮件"
}

func (e *email) Update(message string) {
	fmt.Printf("通过【%s】发送消息:%s\n", e.Name(), message)
}

// telephone 信用卡消息电话订阅者
type telephone struct{}

func (t *telephone) Name() string {
	return "电话"
}

func (t *telephone) Update(message string) {
	fmt.Printf("通过【%s】告知:%s\n", t.Name(), message)
}
