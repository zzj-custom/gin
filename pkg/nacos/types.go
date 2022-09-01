package nacos

type Config struct {
	Host        string `toml:"host"`
	Port        uint64 `toml:"port"`
	Scheme      string `toml:"scheme"`
	ContextPath string `toml:"context_path"`
	NamespaceId string `toml:"namespace_id"`
	UserName    string `toml:"user_name"`
	Password    string `toml:"password"`
	TimeoutMS   uint64 `toml:"timeout_ms"`
	PageSize    uint64 `toml:"page_size"` // 查询每页显示多少数据
}

type Handler interface {
	Handle(dataId string, content string)
}
