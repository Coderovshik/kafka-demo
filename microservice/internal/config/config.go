package config

import (
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Kafka      KafkaConfig
	HTTPServer HTTPServerConfig
}

func New() *Config {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

type KafkaConfig struct {
	SourceTopic string `env:"KAFKA_SOURCE_TOPIC" env-required:"true"`
	TargetTopic string `env:"KAFKA_TARGET_TOPIC" env-required:"true"`
	Host        string `env:"KAFKA_HOST" env-required:"true"`
	Port        string `env:"KAFKA_PORT" env-required:"true"`
}

func (k KafkaConfig) Address() string {
	return net.JoinHostPort(k.Host, k.Port)
}

type HTTPServerConfig struct {
	Host string `env:"HTTP_HOST" env-required:"true"`
	Port string `env:"HTTP_PORT" env-required:"true"`
}

func (h HTTPServerConfig) Address() string {
	return net.JoinHostPort(h.Host, h.Port)
}
