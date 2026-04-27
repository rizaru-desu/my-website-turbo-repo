package config

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type AuthConfig struct {
	Secret                      string
	Issuer                      string
	JWTSecret                   string
	AccessTokenTTL              time.Duration
	RememberMeTTL               time.Duration
	EmailVerificationCooldown   time.Duration
	EmailVerificationMaxPerHour int
	AuthRouteRateLimitWindow    time.Duration
	AuthRouteRateLimitMax       int
	Cookie                      AuthCookieConfig
}

type AuthCookieConfig struct {
	Name     string
	Path     string
	Domain   string
	Secure   bool
	HTTPOnly bool
	SameSite http.SameSite
}

func LoadAuthConfig(environment string) AuthConfig {
	secureDefault := !strings.EqualFold(strings.TrimSpace(environment), "development")
	secret := firstEnv("AUTH_SECRET", "BETTER_AUTH_SECRET", "AUTH_JWT_SECRET")
	if secret == "" {
		secret = "dev-only-change-this-auth-secret"
	}

	return AuthConfig{
		Secret:                      secret,
		Issuer:                      stringEnv("AUTH_JWT_ISSUER", "portfolio-lightweight"),
		JWTSecret:                   firstNonEmptyEnv(secret, "AUTH_JWT_SECRET"),
		AccessTokenTTL:              durationEnv("AUTH_JWT_EXPIRES_IN", 24*time.Hour),
		RememberMeTTL:               durationEnv("AUTH_REMEMBER_ME_EXPIRES_IN", 30*24*time.Hour),
		EmailVerificationCooldown:   durationEnv("AUTH_EMAIL_VERIFICATION_COOLDOWN", time.Minute),
		EmailVerificationMaxPerHour: intEnv("AUTH_EMAIL_VERIFICATION_MAX_PER_HOUR", 5),
		AuthRouteRateLimitWindow:    durationEnv("AUTH_AUTH_ROUTE_RATE_LIMIT_WINDOW", time.Minute),
		AuthRouteRateLimitMax:       intEnv("AUTH_AUTH_ROUTE_RATE_LIMIT_MAX", 5),
		Cookie: AuthCookieConfig{
			Name:     stringEnv("AUTH_COOKIE_NAME", "portfolio_auth"),
			Path:     stringEnv("AUTH_COOKIE_PATH", "/"),
			Domain:   strings.TrimSpace(os.Getenv("AUTH_COOKIE_DOMAIN")),
			Secure:   boolEnv("AUTH_COOKIE_SECURE", secureDefault),
			HTTPOnly: boolEnv("AUTH_COOKIE_HTTP_ONLY", true),
			SameSite: sameSiteEnv("AUTH_COOKIE_SAME_SITE", http.SameSiteLaxMode),
		},
	}
}

func firstNonEmptyEnv(fallback string, keys ...string) string {
	for _, key := range keys {
		if value := strings.TrimSpace(os.Getenv(key)); value != "" {
			return value
		}
	}

	return fallback
}

func stringEnv(key string, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}

	return fallback
}

func durationEnv(key string, fallback time.Duration) time.Duration {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	duration, err := time.ParseDuration(value)
	if err == nil && duration > 0 {
		return duration
	}

	seconds, err := strconv.Atoi(value)
	if err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second
	}

	return fallback
}

func intEnv(key string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return fallback
	}

	return parsed
}

func sameSiteEnv(key string, fallback http.SameSite) http.SameSite {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(key))) {
	case "strict":
		return http.SameSiteStrictMode
	case "none":
		return http.SameSiteNoneMode
	case "lax":
		return http.SameSiteLaxMode
	default:
		return fallback
	}
}
