package services

import (
	"github.com/graphlog/heimdall/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewAMQPConnection(c *config.Config) *amqp.Connection {
	connection, err := amqp.Dial(c.DataStores.AMQPURI)

	if err != nil {
		logrus.Panic("Cannot connect to RabbitMQ")
		return nil
	}

	return connection
}
