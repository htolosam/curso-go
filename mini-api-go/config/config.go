package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	JwtSecret   string
	DatabaseURL string
}

var AppConfig *Config

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("not found .env file")
	}
	appConfig := &Config{
		Port:        getEnv("PORT", "5050"),
		JwtSecret:   getEnv("JWT_SECRET", "someKey"),
		DatabaseURL: getEnv("DATABASE_URL", "mysql"),
	}
	return appConfig
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
