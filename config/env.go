package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnvWithDefault(key string, defaultVal string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultVal
	}

	return val
}

func DatabaseDSN() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to load .env file: %w", err))
	}

	postgresUser := getEnvWithDefault("POSTGRES_USER", "postgres")
	postgresDatabase := getEnvWithDefault("POSTGRES_DB", postgresUser)

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		getEnvWithDefault("POSTGRES_HOST", "postgres"),
		postgresUser,
		os.Getenv("POSTGRES_PASSWORD"),
		postgresDatabase,
		getEnvWithDefault("POSTGRES_PORT", "5432"),
		getEnvWithDefault("POSTGRES_TIMEZONE", "UTC"),
	)
}
