package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	Port              string `json:"port"`
	QueueSize         int    `json:"queue_size"`
	WorkerCount       int    `json:"worker_count"`
	BaseDelayMS       int    `json:"base_delay_ms"`
	PerRequestDelayMS int    `json:"per_request_delay_ms"`
	BatchSize         int    `json:"batch_size"`
	BatchTimeoutMS    int    `json:"batch_timeout_ms"`
}

// Load initializes the configuration from environment variables or defaults.
func Load() *Config {
	// 1. Load defaults and .env
	if err := godotenv.Load(); err != nil {
		slog.Info("no .env file found, relying on system environment variables")
	}

	cfg := &Config{
		Port:              getEnv("PORT", ":8080"),
		QueueSize:         getEnvAsInt("QUEUE_SIZE", 100),
		WorkerCount:       getEnvAsInt("WORKER_COUNT", 4),
		BaseDelayMS:       getEnvAsInt("BASE_DELAY_MS", 100),
		PerRequestDelayMS: getEnvAsInt("PER_REQUEST_DELAY_MS", 50),
		BatchSize:         getEnvAsInt("BATCH_SIZE", 5),
		BatchTimeoutMS:    getEnvAsInt("BATCH_TIMEOUT_MS", 100),
	}

	// 2. Layer scenario if provided
	scenario := os.Getenv("SCENARIO")
	if scenario != "" {
		scenarioPath := filepath.Join("configs", "perf", fmt.Sprintf("%s.json", scenario))
		slog.Info("applying performance scenario", "scenario", scenario, "path", scenarioPath)

		data, err := os.ReadFile(scenarioPath)
		if err != nil {
			slog.Error("failed to read scenario file", "path", scenarioPath, "error", err)
			return cfg
		}

		if err := json.Unmarshal(data, cfg); err != nil {
			slog.Error("failed to unmarshal scenario JSON", "path", scenarioPath, "error", err)
		}
	}

	return cfg
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
