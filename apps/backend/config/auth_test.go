package config

import "testing"

func TestLoadAuthConfigUsesAuthSecretAsPrimarySecret(t *testing.T) {
	t.Setenv("AUTH_SECRET", "legacy-secret")
	t.Setenv("AUTH_JWT_SECRET", "")

	cfg := LoadAuthConfig("development")

	if cfg.Secret != "legacy-secret" {
		t.Fatalf("expected auth secret to be loaded, got %q", cfg.Secret)
	}
	if cfg.JWTSecret != "legacy-secret" {
		t.Fatalf("expected jwt secret to fall back to auth secret, got %q", cfg.JWTSecret)
	}
}

func TestLoadAuthConfigAllowsDedicatedJWTSecret(t *testing.T) {
	t.Setenv("AUTH_SECRET", "legacy-secret")
	t.Setenv("AUTH_JWT_SECRET", "jwt-secret")

	cfg := LoadAuthConfig("development")

	if cfg.Secret != "legacy-secret" {
		t.Fatalf("expected auth secret to stay primary, got %q", cfg.Secret)
	}
	if cfg.JWTSecret != "jwt-secret" {
		t.Fatalf("expected dedicated jwt secret, got %q", cfg.JWTSecret)
	}
}
