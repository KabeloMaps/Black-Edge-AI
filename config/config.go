package config

import "os"

type Config struct {
	PostgresDSN string
	MongoURI string
	ServerPort string
}

func LoadConfig() *Config{
	return &Config{
		PostgresDSN: getEnv("POSTGRES_DSN", "host-localhost user=postgres password=postgres dbname=blackedge port=5432 sslmode=disable"),
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		ServerPort: getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string{
	if v := os.Getenv(key); v != ""{
		return v
	}
	return fallback
}