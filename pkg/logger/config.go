package logger

type FileConfig struct {
	FilePath string `toml:"file_path"`
	FileName string `toml:"file_name"`
	LogLevel int    `toml:"log_level"`
}
