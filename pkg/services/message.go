package services

import (
	"encoding/json"

	"github.com/graphlog/heimdall/pkg/config"
	"github.com/graphlog/heimdall/pkg/models"
	"github.com/streadway/amqp"
)

type MessageService struct {
	Connection *amqp.Connection
	QueueName  string
}

func NewMessageService(conf *config.Config, c *amqp.Connection) *MessageService {
	return &MessageService{
		Connection: c,
		QueueName:  conf.DataStores.AMQPConfig.QueueName,
	}
}

func (m *MessageService) SendLogs(log *models.Log) error {
	bJson, err := json.Marshal(log)

	if err != nil {
		// Failed to marshal json
		return err
	}

	channel, err := m.Connection.Channel()
	defer channel.Close()

	if err != nil {
		return err
	}

	err = channel.Publish("", m.QueueName, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         bJson,
	})

	return err
}
