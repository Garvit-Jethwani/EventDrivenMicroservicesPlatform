package config

import (
	"os"
)

type Config struct {
	HTTPAddress   string
	EventStore    string
	KafkaBrokers  []string
	KafkaTopic    string
	NATSURL       string
	NATSClusterID string
}

func LoadConfig() (*Config, error) {
	httpAddress := os.Getenv("HTTP_ADDRESS")
	eventStore := os.Getenv("EVENT_STORE")
	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	natsURL := os.Getenv("NATS_URL")
	natsClusterID := os.Getenv("NATS_CLUSTER_ID")

	// Validate and parse configuration values

	return &Config{
		HTTPAddress:   httpAddress,
		EventStore:    eventStore,
		KafkaBrokers:  parseKafkaBrokers(kafkaBrokers),
		KafkaTopic:    kafkaTopic,
		NATSURL:       natsURL,
		NATSClusterID: natsClusterID,
	}, nil
}

func parseKafkaBrokers(brokersList string) []string {
	// Parse comma-separated Kafka broker list
	// ...
	return nil
}
