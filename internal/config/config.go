package internal

import "github.com/joho/godotenv"

func Load(configPath string) error {
	err := godotenv.Load(configPath)
	if err != nil {
		return err
	}
	return nil
}

type PGConfig interface {
	DSN() string
}

type HTTPConfig interface {
	Adress() string
}

type KafkaConfig interface {
	Brokers() []string
	Topic() string
	GroupID() string
}
