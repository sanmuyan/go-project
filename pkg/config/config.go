package config

type Config struct {
	LogLevel   int    `mapstructure:"log_level"`
	ServerBind string `mapstructure:"server_bind"`
}

var Conf Config
