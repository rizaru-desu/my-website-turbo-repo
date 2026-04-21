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
	monitoringinfra "api/internal/infrastructure/monitoring"
	"api/internal/infrastructure/persistence/ent"
	postgresinfra "api/internal/infrastructure/persistence/postgres"
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

	handler := newHandler()

	log.Println("Server berjalan di http://localhost:3333")
	if err := http.ListenAndServe(":3333", handler); err != nil {
		log.Println("Gagal menjalankan server:", err)
	}
}

func newHandler() http.Handler {
	mux := http.NewServeMux()

	environment := appEnvironment()
	metricsProvider := monitoringinfra.NewLinuxMetricsProvider()
	systemHandler := httpHandler.NewSystemHandler(version, environment, systemusecase.NewHealthService(time.Now(), appStoragePath(), metricsProvider))
	systemHandler.RegisterRoutes(mux)
	if isDevelopmentEnvironment(environment) {
		httpHandler.NewSwaggerHandler(version).RegisterRoutes(mux)
	}

	securityConfig := config.SecurityConfigForEnvironment(environment)
	corsConfig := config.DefaultCORSConfig()

	// Apply CORS before security headers so preflight requests are answered cleanly.
	return middleware.NewCORS(corsConfig)(middleware.NewSecurityHeaders(securityConfig)(mux))
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
