package config

import (
	"strconv"
	"strings"
)

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	FromName string
}

func LoadSMTPConfig() SMTPConfig {
	port := 587
	if raw := strings.TrimSpace(firstEnv("SMTP_PORT")); raw != "" {
		if parsed, err := strconv.Atoi(raw); err == nil && parsed > 0 {
			port = parsed
		}
	}

	return SMTPConfig{
		Host:     firstEnv("SMTP_HOST"),
		Port:     port,
		Username: firstEnv("SMTP_USERNAME", "SMTP_USER"),
		Password: firstEnv("SMTP_PASSWORD", "SMTP_PASS"),
		From:     firstEnv("SMTP_FROM", "SMTP_FROM_EMAIL"),
		FromName: stringEnv("SMTP_FROM_NAME", "Portfolio Admin"),
	}
}

func (c SMTPConfig) Enabled() bool {
	return c.Host != "" && c.From != ""
}
