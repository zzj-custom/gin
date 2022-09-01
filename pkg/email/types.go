package email

type Config struct {
	From     string `toml:"from"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	options  *Options
}

type Receiver struct {
	Email string
}

type Options struct {
	mailTo     []string
	subject    string
	text       []byte
	html       []byte
	carbonCopy []string
	attachFile []string
}

type Option func(opt *Options)
