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

	host := firstEnv("SMTP_HOST")
	from := firstEnv("SMTP_FROM", "SMTP_FROM_EMAIL")
	username := firstEnv("SMTP_USERNAME", "SMTP_USER")
	password := firstEnv("SMTP_PASSWORD", "SMTP_PASS")
	if isGmailSMTPHost(host) {
		username = gmailUsername(username, from)
		password = strings.ReplaceAll(password, " ", "")
	}

	return SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
		FromName: stringEnv("SMTP_FROM_NAME", "Portfolio Admin"),
	}
}

func (c SMTPConfig) Enabled() bool {
	return c.Host != "" && c.From != ""
}

func isGmailSMTPHost(host string) bool {
	host = strings.ToLower(strings.TrimSpace(host))
	return host == "smtp.gmail.com" || strings.HasSuffix(host, ".smtp.gmail.com")
}

func gmailUsername(username string, from string) string {
	username = strings.TrimSpace(username)
	if strings.Contains(username, "@") {
		return username
	}

	from = strings.TrimSpace(from)
	if strings.Contains(from, "@") {
		return from
	}

	return username
}
