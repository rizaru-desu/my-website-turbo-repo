package config

import "testing"

func TestLoadSMTPConfigNormalizesGmailCredentials(t *testing.T) {
	t.Setenv("SMTP_HOST", "smtp.gmail.com")
	t.Setenv("SMTP_PORT", "587")
	t.Setenv("SMTP_USERNAME", "no-reply")
	t.Setenv("SMTP_PASSWORD", "abcd efgh ijkl mnop")
	t.Setenv("SMTP_FROM", "no-reply@example.com")

	cfg := LoadSMTPConfig()

	if cfg.Username != "no-reply@example.com" {
		t.Fatalf("expected gmail username to fall back to full from email, got %q", cfg.Username)
	}
	if cfg.Password != "abcdefghijklmnop" {
		t.Fatalf("expected gmail app password spaces to be stripped, got %q", cfg.Password)
	}
}

func TestLoadSMTPConfigKeepsFullGmailUsername(t *testing.T) {
	t.Setenv("SMTP_HOST", "smtp.gmail.com")
	t.Setenv("SMTP_USERNAME", "admin@example.com")
	t.Setenv("SMTP_PASSWORD", "abcd efgh")
	t.Setenv("SMTP_FROM", "no-reply@example.com")

	cfg := LoadSMTPConfig()

	if cfg.Username != "admin@example.com" {
		t.Fatalf("expected full gmail username to be preserved, got %q", cfg.Username)
	}
	if cfg.Password != "abcdefgh" {
		t.Fatalf("expected gmail app password spaces to be stripped, got %q", cfg.Password)
	}
}
