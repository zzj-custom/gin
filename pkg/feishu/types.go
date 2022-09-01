package feishu

import "github.com/sirupsen/logrus"

type Config struct {
	Url         string `toml:"url"`
	WebHookUuid string `toml:"web_hook_uuid"`
	Secret      string `toml:"secret"`
	Level       string `toml:"level"`
	AppName     string `toml:"app_name"`
}

type FsHook struct {
	Url         string
	WebHookUuid string
	Secret      string
	Level       []logrus.Level
	AppName     string
}

type FsMsg struct {
	MsgType   string `json:"msg_type"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Content   struct {
		Text string
	} `json:"content"`
}
