package handler

import (
	systemusecase "api/internal/usecase/system"
	"net/http"
)

const apiBasePath = "/api/v1"
const healthPath = "/api/v1/health"

type RootResponse struct {
	Status string `json:"status"`
}

type APIIndexResponse struct {
	Name        string              `json:"name"`
	Version     string              `json:"version"`
	Environment string              `json:"environment"`
	Status      string              `json:"status"`
	BasePath    string              `json:"base_path"`
	Description string              `json:"description"`
	Endpoints   []APIEndpointDetail `json:"endpoints"`
}

type APIEndpointDetail struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Protected   bool   `json:"protected"`
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
	mux.HandleFunc("GET "+apiBasePath, h.APIIndex)
	mux.HandleFunc("GET "+apiBasePath+"/{$}", h.APIIndex)
	mux.HandleFunc("GET "+healthPath, h.Health)
}

// Root godoc
func (h *SystemHandler) Root(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, RootResponse{Status: "ok"})
}

// APIIndex godoc
func (h *SystemHandler) APIIndex(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, APIIndexResponse{
		Name:        "Portfolio Lightweight API",
		Version:     h.version,
		Environment: h.environment,
		Status:      "operational",
		BasePath:    apiBasePath,
		Description: "Backend service for the portfolio platform and admin workspace.",
		Endpoints: []APIEndpointDetail{
			{
				Method:      http.MethodGet,
				Path:        healthPath,
				Description: "Service health and host resource snapshot.",
				Protected:   false,
			},
			{
				Method:      http.MethodPost,
				Path:        "/api/v1/auth/login",
				Description: "Authenticate with email and password, then issue a JWT cookie session.",
				Protected:   false,
			},
			{
				Method:      http.MethodPost,
				Path:        "/api/v1/auth/logout",
				Description: "Revoke the current session token and clear the auth cookie.",
				Protected:   true,
			},
			{
				Method:      http.MethodGet,
				Path:        "/api/v1/auth/me",
				Description: "Return the authenticated user and session claims.",
				Protected:   true,
			},
			{
				Method:      http.MethodGet,
				Path:        "/api/v1/auth/options",
				Description: "Expose safe authentication configuration for clients.",
				Protected:   false,
			},
		},
	})
}

// Health godoc
func (h *SystemHandler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, HealthResponse{
		Version:        h.version,
		HealthSnapshot: h.healthService.Snapshot(),
	})
}
