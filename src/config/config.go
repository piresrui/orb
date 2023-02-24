package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// EnvConfig contains environment variables
type EnvConfig struct {
	Hostname     string
	AssetDir     string
	SignupPath   string
	ReportPath   string
	SignupPeriod int
	ReportPeriod int
}

func ProvideConfig() (*EnvConfig, error) {
	host := GetOrDefault("API_HOST", "localhost")
	port := GetIntOrDefault("API_PORT", 1080)
	assetDir := GetOrDefault("ASSET_DIR", "./assets")
	signupPath := GetOrDefault("SIGNUP_PATH", "/signup")
	reportPath := GetOrDefault("REPORT_PATH", "/status")
	signup := GetIntOrDefault("SIGNUP_PERIOD", 10)
	status := GetIntOrDefault("REPORT_PERIOD", 10)

	return &EnvConfig{
		Hostname:     fmt.Sprintf("http://%s:%d", host, port),
		AssetDir:     assetDir,
		SignupPath:   signupPath,
		ReportPath:   reportPath,
		SignupPeriod: signup,
		ReportPeriod: status,
	}, nil
}

func GetOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Printf("WARN: %s not set, defaulting to %s\n", key, fallback)
		return fallback
	}
	return value
}

func GetIntOrDefault(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Printf("WARN: %s not set, defaulting to %d\n", key, fallback)
		return fallback
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
