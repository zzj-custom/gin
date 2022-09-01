package email_event

import (
	"go-api/event/dispatcher"
)

type EmailEventInterface interface {
	dispatcher.Evt
}

type EmailInfo struct {
	MailTo     []string
	Subject    string
	Text       string
	Html       string
	CarbonCopy []string
	AttachFile []string
}

type EmailEvent struct {
	AoEvents chan *EmailInfo
}
