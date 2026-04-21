package handler

import (
	systemusecase "api/internal/usecase/system"
	"net/http"
)

const healthPath = "/api/v1/health"

type RootResponse struct {
	Status string `json:"status"`
}

type HealthResponse struct {
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

// Root godoc
//
//	@Summary		Get root status
//	@Description	Returns a minimal root status response without exposing implementation details.
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	RootResponse
//	@Router			/ [get]
func (h *SystemHandler) Root(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, RootResponse{Status: "ok"})
}

// Health godoc
//
//	@Summary		Get service health
//	@Description	Returns application and host health information.
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	HealthResponse
//	@Router			/api/v1/health [get]
func (h *SystemHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, HealthResponse{
		Version:        h.version,
		HealthSnapshot: h.healthService.Snapshot(),
	})
}
