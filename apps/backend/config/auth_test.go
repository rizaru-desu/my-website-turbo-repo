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

func TestLoadAuthConfigLoadsEmailVerificationAndRateLimitSettings(t *testing.T) {
	t.Setenv("AUTH_EMAIL_VERIFICATION_COOLDOWN", "2m")
	t.Setenv("AUTH_EMAIL_VERIFICATION_MAX_PER_HOUR", "7")
	t.Setenv("AUTH_AUTH_ROUTE_RATE_LIMIT_WINDOW", "30s")
	t.Setenv("AUTH_AUTH_ROUTE_RATE_LIMIT_MAX", "3")

	cfg := LoadAuthConfig("development")

	if cfg.EmailVerificationCooldown.String() != "2m0s" {
		t.Fatalf("expected email verification cooldown 2m, got %s", cfg.EmailVerificationCooldown)
	}
	if cfg.EmailVerificationMaxPerHour != 7 {
		t.Fatalf("expected email verification max per hour 7, got %d", cfg.EmailVerificationMaxPerHour)
	}
	if cfg.AuthRouteRateLimitWindow.String() != "30s" {
		t.Fatalf("expected auth route rate limit window 30s, got %s", cfg.AuthRouteRateLimitWindow)
	}
	if cfg.AuthRouteRateLimitMax != 3 {
		t.Fatalf("expected auth route rate limit max 3, got %d", cfg.AuthRouteRateLimitMax)
	}
}
