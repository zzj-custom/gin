package dispatcher

import (
	"context"
	"github.com/ethereum/go-ethereum/event"
)

type ManagerEvent struct {
	feeds map[string]*event.Feed
}

type Evt interface {
	Name() string
}

type Task interface {
	Run(ctx context.Context)
}
