package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"api/config"
	httpHandler "api/internal/delivery/http/handler"
	"api/internal/delivery/http/middleware"
	authinfra "api/internal/infrastructure/auth"
	mailinfra "api/internal/infrastructure/mail"
	monitoringinfra "api/internal/infrastructure/monitoring"
	authstore "api/internal/infrastructure/persistence/auth"
	"api/internal/infrastructure/persistence/ent"
	postgresinfra "api/internal/infrastructure/persistence/postgres"
	authusecase "api/internal/usecase/auth"
	systemusecase "api/internal/usecase/system"
)

var version = "1.0.0"

func main() {
	if err := config.LoadAppEnv(); err != nil {
		log.Fatal("Gagal memuat konfigurasi environment:", err)
	}

	databaseClient, err := newDatabaseClient(context.Background())
	if err != nil {
		log.Fatal("Gagal menghubungkan PostgreSQL:", err)
	}
	if databaseClient != nil {
		defer databaseClient.Close()
	}

	handler := newHandler(databaseClient)

	log.Println("Server berjalan di http://localhost:3333")
	if err := http.ListenAndServe(":3333", handler); err != nil {
		log.Println("Gagal menjalankan server:", err)
	}
}

func newHandler(databaseClients ...*ent.Client) http.Handler {
	mux := http.NewServeMux()

	environment := appEnvironment()
	authConfig := config.LoadAuthConfig(environment)
	databaseClient := optionalDatabaseClient(databaseClients...)

	metricsProvider := monitoringinfra.NewLinuxMetricsProvider()
	systemHandler := httpHandler.NewSystemHandler(version, environment, systemusecase.NewHealthService(time.Now(), appStoragePath(), metricsProvider))
	systemHandler.RegisterRoutes(mux)
	authHandler := httpHandler.NewAuthHandler(
		newAuthService(databaseClient, authConfig),
		httpHandler.AuthCookieConfig{
			Name:     authConfig.Cookie.Name,
			Path:     authConfig.Cookie.Path,
			Domain:   authConfig.Cookie.Domain,
			Secure:   authConfig.Cookie.Secure,
			HTTPOnly: authConfig.Cookie.HTTPOnly,
			SameSite: authConfig.Cookie.SameSite,
		},
	)
	authHandler.RegisterRoutes(mux)

	securityConfig := config.SecurityConfigForEnvironment(environment)
	corsConfig := config.DefaultCORSConfig()
	rateLimitConfig := middleware.IPRateLimitConfig{
		Window: authConfig.AuthRouteRateLimitWindow,
		Max:    authConfig.AuthRouteRateLimitMax,
		Routes: map[string]struct{}{
			"POST /api/v1/auth/email-verification": {},
			"POST /api/v1/auth/forgot-password":    {},
			"POST /api/v1/auth/reset-password":     {},
		},
	}

	// Apply CORS before security headers so preflight requests are answered cleanly.
	return middleware.NewCORS(corsConfig)(middleware.NewSecurityHeaders(securityConfig)(middleware.NewIPRateLimiter(rateLimitConfig)(mux)))
}

func newAuthService(databaseClient *ent.Client, authConfig config.AuthConfig) *authusecase.Service {
	clock := authinfra.SystemClock{}
	smtpConfig := config.LoadSMTPConfig()

	svc := authusecase.NewService(
		authstore.NewEntCredentialRepository(databaseClient),
		authstore.NewEntSessionRepository(databaseClient, clock.Now),
		authinfra.NewJWTManager(authConfig.JWTSecret, authConfig.Issuer),
		authinfra.NewPasswordVerifier(authConfig.Secret),
		authinfra.CryptoIDGenerator{},
		clock,
		authusecase.Options{
			Issuer:                    authConfig.Issuer,
			AccessTokenTTL:            authConfig.AccessTokenTTL,
			RememberMeTTL:             authConfig.RememberMeTTL,
			EmailVerificationCooldown: authConfig.EmailVerificationCooldown,
			EmailVerificationMax:      authConfig.EmailVerificationMaxPerHour,
			FrontendURL:               stringEnvFallback("FRONTEND_URL", "http://localhost:3111"),
			BackendURL:                stringEnvFallback("BACKEND_URL", "http://localhost:3333"),
		},
	)

	if databaseClient != nil {
		svc.SetVerifications(authstore.NewEntVerificationRepository(databaseClient, clock.Now))
		svc.SetEmailVerificationLimiter(authstore.NewEntEmailVerificationLimiter(databaseClient))
		svc.SetTwoFactors(authstore.NewEntTwoFactorRepository(databaseClient))
	}

	totpIssuer := stringEnvFallback("TOTP_ISSUER", "PortfolioAdmin")
	svc.SetTOTP(authinfra.NewTOTPManager(totpIssuer, authConfig.Secret))
	log.Println("TOTP 2FA dikonfigurasi, issuer:", totpIssuer)

	if smtpConfig.Enabled() {
		svc.SetMailSender(mailinfra.NewSMTPMailSender(smtpConfig))
		log.Println("SMTP mail sender dikonfigurasi:", smtpConfig.Host)
	} else {
		log.Println("SMTP belum dikonfigurasi; forgot password dan email verification tidak aktif")
	}

	return svc
}

func stringEnvFallback(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return strings.TrimRight(value, "/")
	}

	return fallback
}

func optionalDatabaseClient(databaseClients ...*ent.Client) *ent.Client {
	if len(databaseClients) == 0 {
		return nil
	}

	return databaseClients[0]
}

func newDatabaseClient(ctx context.Context) (*ent.Client, error) {
	databaseConfig := config.LoadDatabaseConfig()
	if !databaseConfig.Enabled() {
		log.Println("DATABASE_URL belum diset; koneksi PostgreSQL dilewati")
		return nil, nil
	}

	client, err := postgresinfra.NewClient(ctx, databaseConfig)
	if err != nil {
		return nil, fmt.Errorf("bootstrap database: %w", err)
	}

	if databaseConfig.AutoMigrate {
		log.Println("PostgreSQL terhubung dan migrasi Ent sudah dijalankan")
	} else {
		log.Println("PostgreSQL terhubung")
	}

	return client, nil
}

func appEnvironment() string {
	if value := os.Getenv("APP_ENV"); value != "" {
		return value
	}

	return "development"
}

func isDevelopmentEnvironment(environment string) bool {
	return strings.EqualFold(strings.TrimSpace(environment), "development")
}

func appStoragePath() string {
	workingDir, err := os.Getwd()
	if err != nil {
		return "."
	}

	return workingDir
}
