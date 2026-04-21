package config

import (
	"os"
	"strings"
)

type DatabaseConfig struct {
	URL         string
	AutoMigrate bool
}

func LoadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		URL:         firstEnv("DATABASE_URL", "POSTGRES_URL", "POSTGRES_DATABASE_URL"),
		AutoMigrate: boolEnv("DB_AUTO_MIGRATE", false),
	}
}

func (c DatabaseConfig) Enabled() bool {
	return strings.TrimSpace(c.URL) != ""
}

func firstEnv(keys ...string) string {
	for _, key := range keys {
		if value := strings.TrimSpace(os.Getenv(key)); value != "" {
			return value
		}
	}

	return ""
}

func boolEnv(key string, fallback bool) bool {
	value := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	switch value {
	case "1", "true", "yes", "y", "on":
		return true
	case "0", "false", "no", "n", "off":
		return false
	default:
		return fallback
	}
}
