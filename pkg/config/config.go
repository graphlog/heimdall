package config

import "os"

type ServerConfig struct {
	Port string
}

type Config struct {
	Server *ServerConfig
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Port: envOrDefault("PORT", "7000"),
		},
	}
}

func envOrDefault(name string, d string) string {
	envVar := os.Getenv(name)
	if len(envVar) > 0 {
		return envVar
	}

	return d
}
