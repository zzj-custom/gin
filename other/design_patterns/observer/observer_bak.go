package observer

type Event struct {
	Data int
}

type ObserverInterface interface {
	// Update 更新事件
	Update(*Event)
}

type SubjectInterface interface {
	// Register 注册观察者
	Register(observer ObserverInterface)

	// Delete 删除观察者
	Delete(observer ObserverInterface)

	// Notify 通知观察者
	Notify(event *Event)
}

type Observer struct {
	Account string
}

func (co *Observer) Update(event *Event) {
	// 更新数据
	//LandsSubscribe(event.Data)
}

type Subject struct {
	// map方便删除，定义切片也是可以，[]Observer
	Observers map[ObserverInterface]struct{}
}

func (cs *Subject) Register(observer ObserverInterface) {
	cs.Observers[observer] = struct{}{}
}

func (cs *Subject) Delete(observer ObserverInterface) {
	// delete 直接删除map里面的某个值
	delete(cs.Observers, observer)
}

func (cs *Subject) Notify(event *Event) {
	for observer := range cs.Observers {
		observer.Update(event)
	}
}

func Init() *Subject {
	return &Subject{
		Observers: make(map[ObserverInterface]struct{}),
	}
}
