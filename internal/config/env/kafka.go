package env

import (
	"errors"
	"os"
	"strings"
	config "wb/internal/config"
)

var _ config.KafkaConfig = (*kafkaConfig)(nil)

const (
	kafkaBrokersEnvName = "KAFKA_BROKERS"
	kafkaTopicEnvName   = "KAFKA_TOPIC"
	kafkaGroupIDEnvName = "KAFKA_GROUP_ID"
)

type kafkaConfig struct {
	brokers []string
	topic   string
	groupID string
}

func NewKafkaConfig() (*kafkaConfig, error) {
	brokers := os.Getenv(kafkaBrokersEnvName)
	if len(brokers) == 0 {
		return nil, errors.New("kafka brokers not found")
	}
	topic := os.Getenv(kafkaTopicEnvName)
	if len(topic) == 0 {
		return nil, errors.New("kafka topic not found")
	}
	groupID := os.Getenv(kafkaGroupIDEnvName)
	if len(groupID) == 0 {
		return nil, errors.New("kafka group id not found")
	}
	return &kafkaConfig{
		brokers: strings.Split(brokers, ","),
		topic:   topic,
		groupID: groupID,
	}, nil
}

func (cfg *kafkaConfig) Brokers() []string {
	return cfg.brokers
}

func (cfg *kafkaConfig) Topic() string {
	return cfg.topic
}

func (cfg *kafkaConfig) GroupID() string {
	return cfg.groupID
}
