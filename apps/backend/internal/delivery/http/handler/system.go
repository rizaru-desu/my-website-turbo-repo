package handler

import (
	systemusecase "api/internal/usecase/system"
	"net/http"
)

const healthPath = "/api/v1/health"

type rootResponse struct {
	Status string `json:"status"`
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
	writeJSON(w, http.StatusOK, rootResponse{Status: "ok"})
}

func (h *SystemHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Version:        h.version,
		HealthSnapshot: h.healthService.Snapshot(),
	})
}
