package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Port             string
	QueueSize        int
	WorkerCount      int
	InferenceDelayMS int
	BatchSize        int
}

// Load initializes the configuration from environment variables or defaults.
func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		slog.Info("no .env file found, relying on system environment variables")
	}

	return &Config{
		Port:             getEnv("PORT", ":8080"),
		QueueSize:        getEnvAsInt("QUEUE_SIZE", 100),
		WorkerCount:      getEnvAsInt("WORKER_COUNT", 4),
		InferenceDelayMS: getEnvAsInt("INFERENCE_DELAY_MS", 500),
		BatchSize:        getEnvAsInt("BATCH_SIZE", 5),
	}
}

// Helper to get string env variable with a default
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		// If PORT is just a number like "8080", append ":"
		if key == "PORT" && value[0] != ':' {
			return ":" + value
		}
		return value
	}
	return defaultValue
}

// Helper to get int env variable with a default
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}
