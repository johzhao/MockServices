package config

type Config struct {
	Logger Logger `yaml:"logger"`
}

type Logger struct {
	Level     string `yaml:"level"`
	Filepath  string `yaml:"filepath"`
	MaxSize   int    `yaml:"max_size"`
	MaxBackup int    `yaml:"max_backup"`
	MaxAge    int    `yaml:"max_age"`
	Compress  bool   `yaml:"compress"`
}

type Server struct {
	Address string `yaml:"address"`
}
