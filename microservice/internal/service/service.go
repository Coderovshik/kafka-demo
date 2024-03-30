package service

import (
	"errors"
	"log"

	"github.com/Coderovshik/kafka-demo/microservice/internal/config"
	"github.com/Coderovshik/kafka-demo/microservice/internal/domain"
	"github.com/IBM/sarama"
)

var (
	ErrConsumerChannelClosed = errors.New("channel closed")
	ErrTimedOut              = errors.New("timed out")
)

type Service struct {
	cfg       *config.Config
	messageCh chan *domain.Message
	producer  sarama.SyncProducer
	consumer  sarama.Consumer
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
		cfg:       cfg,
		messageCh: make(chan *domain.Message),
		producer:  producer,
		consumer:  consumer,
	}
}

func (s *Service) Consume() error {
	partConsumer, err := s.consumer.ConsumePartition(s.cfg.Kafka.SourceTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
		return err
	}
	defer partConsumer.Close()

	for {
		msg, ok := <-partConsumer.Messages()
		if !ok {
			log.Println("ERROR: channel closed")
			return ErrConsumerChannelClosed
		}

		message, err := ToDomainMessage(msg)
		if err != nil {
			log.Println("ERROR: failed conversion to domain message")
			return err
		}
		log.Printf("INFO: received message: %+v\n", message)

		s.messageCh <- message
	}
}

func (s *Service) Produce() error {
	for {
		msg := <-s.messageCh
		prodMsg, err := NewProducerMessage(msg, s.cfg.Kafka.TargetTopic)
		if err != nil {
			log.Println("ERROR: failed conversion to producer message")
			return err
		}

		partition, offset, err := s.producer.SendMessage(prodMsg)
		if err != nil {
			log.Println("ERROR: failed to send producer message")
			return err
		}
		log.Printf("Message sent: partition: %d, offset: %d\n", partition, offset)
	}
}
