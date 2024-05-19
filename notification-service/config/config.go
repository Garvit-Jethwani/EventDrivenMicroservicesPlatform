package config

import (
	"os"
)

type Config struct {
	EventStore       string
	EventStoreConfig map[string]string
}

func LoadConfig() (*Config, error) {
	eventStore := os.Getenv("EVENT_STORE")
	eventStoreConfig := make(map[string]string)
	// Load event store configuration from environment variables

	return &Config{
		EventStore:       eventStore,
		EventStoreConfig: eventStoreConfig,
	}, nil
}
