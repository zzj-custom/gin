package mysql

type Database struct {
	DSN             string `json:"dsn" toml:"dsn"`
	Username        string `json:"username" toml:"username"`
	Password        string `json:"password" toml:"password"`
	Host            string `json:"host" toml:"host"`
	Port            int    `json:"port" toml:"port"`
	Database        string `json:"database" toml:"database"`
	MaxOpenConn     int    `json:"max_open_conn" toml:"max_open_conn"`
	MaxIdleConn     int    `json:"max_idle_conn" toml:"max_idle_conn"`
	ConnMaxFreeTime int    `json:"conn_max_free_time" toml:"conn_max_free_time"`
	ConnMaxLifeTime int    `json:"conn_max_life_time" toml:"conn_max_life_time"`
	UseLog          bool   `json:"use_log" toml:"use_log"`
	LogLevel        int    `json:"log_level" toml:"log_level"`
	ModelPath       string `json:"model_path" toml:"model_path"`
}
