package config

import "os"

type ServerConfig struct {
	Port string
}

type Config struct {
	Server     *ServerConfig
	DataStores *DS
}

type AMQPConfig struct {
	URI       string
	QueueName string
}

type DS struct {
	AMQPConfig    *AMQPConfig
	ClickHouseURI string
}

func NewConfig() *Config {
	return &Config{
		Server: &ServerConfig{
			Port: envOrDefault("PORT", "7000"),
		},
		DataStores: &DS{
			AMQPConfig: &AMQPConfig{
				URI:       envOrDefault("AMQP_URI", "amqp://localhost:5672"),
				QueueName: envOrDefault("GRAPHLOG_QUEUE_NAME", "graphlog_logger"),
			},
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
