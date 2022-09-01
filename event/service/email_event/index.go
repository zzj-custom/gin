package email_event

import (
	"context"
	"fmt"
	"go-api/config"
	"go-api/event/dispatcher"
	"go-api/internal/task"
	"go-api/pkg/email"
)

func (e EmailEvent) Run(ctx context.Context) {
	e.listen(ctx)
}

func (e EmailEvent) listen(ctx context.Context) {
	fmt.Println("开启EmailEvent事件监听")
	feed := dispatcher.NewManagerEvent().RegisterFeed(Email)
	sub := feed.Subscribe(e.AoEvents)
	defer sub.Unsubscribe()
	taskPool := task.Pool()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("邮件事件监听协程收到退出消息")
			return
		case aoEvt := <-e.AoEvents:
			fmt.Printf("接收到数据，account:%v", aoEvt)
			_ = taskPool.Submit(func() {
				_ = config.Config().Email.SendMail(
					aoEvt.GetMailTo(),
					aoEvt.GetSubject(),
					email.WithTextOptions(aoEvt.Text),
				)
			})
		}
	}
}
