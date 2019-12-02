package services

import (
	"encoding/json"
	"github.com/graphlog/heimdall/pkg/models"
	"github.com/streadway/amqp"
)

type MessageService struct {
	Connection *amqp.Connection
}

func NewMessageService(c *amqp.Connection) *MessageService {
	return &MessageService{
		Connection: c,
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

	err = channel.Publish("", "graphlog_logger", false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         bJson,
	})

	return err
}
