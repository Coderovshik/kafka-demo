package service

import (
	"encoding/json"

	"github.com/Coderovshik/kafka-demo/microservice/internal/domain"
	"github.com/IBM/sarama"
)

func NewProducerMessage(msg *domain.Message, topic string) (*sarama.ProducerMessage, error) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(msg.ID),
		Value: sarama.ByteEncoder(bytes),
	}, nil
}

func ToDomainMessage(msg *sarama.ConsumerMessage) (*domain.Message, error) {
	var message domain.Message
	err := json.Unmarshal(msg.Value, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
