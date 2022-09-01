package service

import (
	"context"
	"go-api/event/service/email_event"
)

func Start(ctx context.Context) {
	emailEvent := new(email_event.EmailEvent)
	emailEvent.AoEvents = make(chan *email_event.EmailInfo, 30)
	emailEvent.Run(ctx)
}
