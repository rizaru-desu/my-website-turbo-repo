package handler

import (
	systemusecase "api/internal/usecase/system"
	"net/http"
	"time"
)

const healthPath = "/api/v1/health"

type serviceMetadata struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Status      string            `json:"status"`
	Environment string            `json:"environment"`
	Description string            `json:"description"`
	Timestamp   string            `json:"timestamp"`
	Endpoints   map[string]string `json:"endpoints"`
}

type healthResponse struct {
	Version string `json:"version"`
	systemusecase.HealthSnapshot
}

type SystemHandler struct {
	version       string
	environment   string
	healthService systemusecase.HealthReporter
}

func NewSystemHandler(version string, environment string, healthService systemusecase.HealthReporter) *SystemHandler {
	return &SystemHandler{
		version:       version,
		environment:   environment,
		healthService: healthService,
	}
}

func (h *SystemHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /{$}", h.Root)
	mux.HandleFunc("GET "+healthPath, h.Health)
}

func (h *SystemHandler) Root(w http.ResponseWriter, r *http.Request) {
	response := serviceMetadata{
		Name:        "Portfolio Lightweight API",
		Version:     h.version,
		Status:      "operational",
		Environment: h.environment,
		Description: "Lightweight backend service for the portfolio platform and admin workspace, optimized for low-resource home lab deployments.",
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
		Endpoints: map[string]string{
			"self":   "/",
			"health": healthPath,
		},
	}

	writeJSON(w, http.StatusOK, response)
}

func (h *SystemHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Version:        h.version,
		HealthSnapshot: h.healthService.Snapshot(),
	})
}
