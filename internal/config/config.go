package config

import "os"

type Config struct {
	AppName string
	AppEnv  string
	AppPort string
}

func Load() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "mining-asset-management"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
