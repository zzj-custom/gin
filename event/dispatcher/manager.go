package dispatcher

import (
	"github.com/ethereum/go-ethereum/event"
	"sync"
)

func (e *ManagerEvent) RegisterFeed(name string) *event.Feed {
	if _, ok := e.feeds[name]; !ok {
		e.feeds[name] = new(event.Feed)
	}
	return e.feeds[name]
}

var (
	managerEvent     *ManagerEvent
	managerEventOnce sync.Once
)

func NewManagerEvent() *ManagerEvent {
	managerEventOnce.Do(func() {
		managerEvent = new(ManagerEvent)
		managerEvent.feeds = make(map[string]*event.Feed)
	})
	return managerEvent
}
