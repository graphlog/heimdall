package config

import "os"

type ServerConfig struct {
	Port string
}

type Config struct {
	Server     *ServerConfig
	DataStores *DS
}

type DS struct {
	AMQPURI       string
	ClickHouseURI string
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Port: envOrDefault("PORT", "7000"),
		},
		DataStores: &DS{
			AMQPURI:       envOrDefault("AMQP_URI", "amqp://localhost:5672"),
			ClickHouseURI: envOrDefault("CLICKHOUSE_URI", "tcp://127.0.0.1:9000?debug=true"),
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
