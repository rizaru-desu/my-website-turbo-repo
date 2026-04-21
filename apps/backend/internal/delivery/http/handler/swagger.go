package handler

import (
	"html/template"
	"net/http"
)

const (
	swaggerPath    = "/swagger/"
	swaggerDocPath = "/swagger/doc.json"
)

type SwaggerHandler struct {
	version string
}

func NewSwaggerHandler(version string) *SwaggerHandler {
	return &SwaggerHandler{version: version}
}

func (h *SwaggerHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /swagger", h.Redirect)
	mux.HandleFunc("GET "+swaggerPath, h.UI)
	mux.HandleFunc("GET "+swaggerDocPath, h.OpenAPISpec)
}

func (h *SwaggerHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, swaggerPath, http.StatusMovedPermanently)
}

func (h *SwaggerHandler) UI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	data := struct {
		SpecURL string
		Title   string
	}{
		SpecURL: swaggerDocPath,
		Title:   "Portfolio Lightweight API Docs",
	}

	if err := swaggerUITemplate.Execute(w, data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *SwaggerHandler) OpenAPISpec(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"openapi": "3.0.3",
		"info": map[string]any{
			"title":       "Portfolio Lightweight API",
			"description": "Lightweight backend service for the portfolio platform and admin workspace.",
			"version":     h.version,
		},
		"servers": []map[string]string{
			{"url": "/"},
		},
		"paths": map[string]any{
			"/": map[string]any{
				"get": map[string]any{
					"summary":     "Get root status",
					"description": "Returns a minimal root status response.",
					"responses": map[string]any{
						"200": map[string]any{
							"description": "Root status",
							"content": map[string]any{
								"application/json": map[string]any{
									"schema": map[string]string{"$ref": "#/components/schemas/RootResponse"},
								},
							},
						},
					},
				},
			},
			healthPath: map[string]any{
				"get": map[string]any{
					"summary":     "Get service health",
					"description": "Returns application and host health information.",
					"responses": map[string]any{
						"200": map[string]any{
							"description": "Health snapshot",
							"content": map[string]any{
								"application/json": map[string]any{
									"schema": map[string]string{"$ref": "#/components/schemas/HealthResponse"},
								},
							},
						},
					},
				},
			},
		},
		"components": map[string]any{
			"schemas": map[string]any{
				"RootResponse": map[string]any{
					"type":     "object",
					"required": []string{"status"},
					"properties": map[string]any{
						"status": map[string]string{"type": "string"},
					},
				},
				"HealthResponse": map[string]any{
					"type":     "object",
					"required": []string{"version", "status", "systemHealth"},
					"properties": map[string]any{
						"version":      map[string]string{"type": "string"},
						"status":       map[string]string{"type": "string"},
						"systemHealth": map[string]string{"$ref": "#/components/schemas/SystemHealth"},
					},
				},
				"SystemHealth": map[string]any{
					"type":     "object",
					"required": []string{"cpuUsage", "memory", "storage", "uptime"},
					"properties": map[string]any{
						"cpuUsage": map[string]string{"type": "string"},
						"memory":   map[string]string{"$ref": "#/components/schemas/ResourceUsage"},
						"storage":  map[string]string{"$ref": "#/components/schemas/StorageUsage"},
						"uptime":   map[string]string{"type": "string"},
					},
				},
				"ResourceUsage": map[string]any{
					"type":     "object",
					"required": []string{"used", "total", "percentage"},
					"properties": map[string]any{
						"used":       map[string]string{"type": "string"},
						"total":      map[string]string{"type": "string"},
						"percentage": map[string]string{"type": "string"},
					},
				},
				"StorageUsage": map[string]any{
					"type":     "object",
					"required": []string{"used", "free", "total", "percentage"},
					"properties": map[string]any{
						"used":       map[string]string{"type": "string"},
						"free":       map[string]string{"type": "string"},
						"total":      map[string]string{"type": "string"},
						"percentage": map[string]string{"type": "string"},
					},
				},
			},
		},
	})
}

var swaggerUITemplate = template.Must(template.New("swagger-ui").Parse(`<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ .Title }}</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css">
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
  <script>
    window.onload = function () {
      window.ui = SwaggerUIBundle({
        url: "{{ .SpecURL }}",
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [
          SwaggerUIBundle.presets.apis
        ],
        layout: "BaseLayout"
      });
    };
  </script>
</body>
</html>`))
