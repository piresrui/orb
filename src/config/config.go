package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// EnvConfig contains environment variables
type EnvConfig struct {
	Hostname   string
	AssetDir   string
	SignupPath string
	ReportPath string
}

func ProvideConfig() (*EnvConfig, error) {
	host, ok := os.LookupEnv("API_HOST")
	if !ok {
		log.Printf("WARN: API_HOST not set")
	}
	portStr, ok := os.LookupEnv("API_PORT")
	if !ok {
		log.Printf("WARN: API_PORT not set")
	}
	assetDir, ok := os.LookupEnv("ASSET_DIR")
	if !ok {
		log.Printf("WARN: ASSET_DIR not set")
	}
	signupPath, ok := os.LookupEnv("SIGNUP_PATH")
	if !ok {
		log.Printf("WARN: SIGNUP_PATH not set")
	}
	reportPath, ok := os.LookupEnv("REPORT_PATH")
	if !ok {
		log.Printf("WARN: REPORT_PATH not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, err
	}

	return &EnvConfig{
		Hostname:   fmt.Sprintf("http://%s:%d/", host, port),
		AssetDir:   assetDir,
		SignupPath: signupPath,
		ReportPath: reportPath,
	}, nil
}
