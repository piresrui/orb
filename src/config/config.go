package config

import (
	"log"
	"os"
	"strconv"
)

// EnvConfig contains environment variables
type EnvConfig struct {
	APIEndpoint string
	APIPort     int
}

func ProvideConfig() (*EnvConfig, error) {
	endpoint, ok := os.LookupEnv("API_ENDPOINT")
	if !ok {
		log.Printf("WARN: API_ENDPOINT not set")
	}
	portStr, ok := os.LookupEnv("API_PORT")
	if !ok {
		log.Printf("WARN: API_PORT not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &EnvConfig{
		APIEndpoint: endpoint,
		APIPort:     port,
	}, nil
}
