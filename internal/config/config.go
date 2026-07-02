package config

import (
	"log"

	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppName string
	AppEnv  string
	AppPort string
	DBHost  string
	DBUser  string
	DBPass  string
	DBName  string
	DBPort  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		AppName: getEnv("APP_NAME", "mining-asset-management"),
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASS", "postgres"),
		DBName:  getEnv("DB_NAME", "mining_asset"),
		DBPort:  getEnv("DB_PORT", "5432"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
