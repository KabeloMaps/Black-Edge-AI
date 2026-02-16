package config

import (
	"os"
	"strings"
)

type Config struct {
	PostgresDSN string
	MongoURI string
	ServerPort string
}

func LoadConfig() *Config{
	return &Config{
		
		PostgresDSN: getEnv(
			"POSTGRES_DSN",
			"host=localhost port=5432 user=stenamatousand dbname=blackedge sslmode=disable",
		),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		ServerPort: strings.TrimSpace(getEnv("SERVER_PORT", "8080")),
	}
}

func getEnv(key, fallback string) string{
	if v := os.Getenv(key); v != ""{
		return v
	}
	return fallback
}
