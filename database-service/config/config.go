package config

import (
	"os"
)

type Config struct {
	GRPCAddress string
	// Database configurations
}

func LoadConfig() (*Config, error) {
	grpcAddress := os.Getenv("GRPC_ADDRESS")
	// Load database configurations from environment variables

	return &Config{
		GRPCAddress: grpcAddress,
		// Set database configurations
	}, nil
}
