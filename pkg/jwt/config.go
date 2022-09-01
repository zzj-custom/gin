package jwt

type Config struct {
	TokenSecret string `toml:"token_secret"`
	AesKey      string `toml:"aes_key"`
	Salt        string `toml:"salt"`
}
