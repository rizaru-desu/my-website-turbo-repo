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

	// Apply CORS before security headers so preflight requests are answered cleanly.
	return middleware.NewCORS(corsConfig)(middleware.NewSecurityHeaders(securityConfig)(mux))
}

func newAuthService(databaseClient *ent.Client, authConfig config.AuthConfig) *authusecase.Service {
	clock := authinfra.SystemClock{}
	return authusecase.NewService(
		authstore.NewEntCredentialRepository(databaseClient),
		authstore.NewEntSessionRepository(databaseClient, clock.Now),
		authinfra.NewJWTManager(authConfig.JWTSecret, authConfig.Issuer),
		authinfra.NewPasswordVerifier(authConfig.Secret),
		authinfra.CryptoIDGenerator{},
		clock,
		authusecase.Options{
			Issuer:         authConfig.Issuer,
			AccessTokenTTL: authConfig.AccessTokenTTL,
			RememberMeTTL:  authConfig.RememberMeTTL,
			Cookie: authusecase.CookieOptions{
				Name:     authConfig.Cookie.Name,
				Secure:   authConfig.Cookie.Secure,
				HTTPOnly: authConfig.Cookie.HTTPOnly,
				SameSite: sameSiteString(authConfig.Cookie.SameSite),
			},
		},
	)
}

func optionalDatabaseClient(databaseClients ...*ent.Client) *ent.Client {
	if len(databaseClients) == 0 {
		return nil
	}

	return databaseClients[0]
}

func sameSiteString(value http.SameSite) string {
	switch value {
	case http.SameSiteStrictMode:
		return "strict"
	case http.SameSiteNoneMode:
		return "none"
	case http.SameSiteLaxMode:
		return "lax"
	default:
		return "default"
	}
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
