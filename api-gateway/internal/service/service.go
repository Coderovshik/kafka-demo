package service

import (
	"context"
	"errors"
	"log"

	"github.com/Coderovshik/kafka-demo/api-gateway/internal/config"
	"github.com/Coderovshik/kafka-demo/api-gateway/internal/domain"
	"github.com/IBM/sarama"
)

var (
	ErrConsumerChannelClosed = errors.New("channel closed")
	ErrTimedOut              = errors.New("timed out")
)

type Service struct {
	cfg      *config.Config
	producer sarama.SyncProducer
	consumer sarama.Consumer
}

func New(cfg *config.Config) *Service {
	producer, err := sarama.NewSyncProducer([]string{cfg.Kafka.Address()}, nil)
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := sarama.NewConsumer([]string{cfg.Kafka.Address()}, nil)
	if err != nil {
		log.Fatal(err)
	}

	return &Service{
		cfg:      cfg,
		producer: producer,
		consumer: consumer,
	}
}

func (s *Service) Produce(ctx context.Context, message *domain.Message) error {
	msg, err := NewProducerMessage(message, s.cfg.Kafka.TargetTopic)
	if err != nil {
		log.Println("ERROR: failed to create producer message")
		return nil
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		log.Println("ERROR: failed to send message to Kafka")
		return err
	}
	log.Printf("Message sent: partition: %d, offset: %d\n", partition, offset)

	return nil
}

func (s *Service) Consume(ctx context.Context) (*domain.Message, error) {
	partConsumer, err := s.consumer.ConsumePartition(s.cfg.Kafka.SourceTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Printf("ERROR: failed to consume partition: %s\n", s.cfg.Kafka.SourceTopic)
		return nil, err
	}
	defer partConsumer.Close()

	select {
	case <-ctx.Done():
		log.Println("INFO: timed out")
		return nil, ErrTimedOut
	case msg, ok := <-partConsumer.Messages():
		if !ok {
			log.Println("Channel closed, exiting gorutine")
			return nil, ErrConsumerChannelClosed
		}

		res, err := ToDomainMessage(msg)
		if err != nil {
			log.Println("ERROR: failed to convert to domain message")
			return nil, err
		}

		return res, nil
	}
}
