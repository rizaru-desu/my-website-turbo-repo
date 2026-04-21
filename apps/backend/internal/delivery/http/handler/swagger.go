package handler

import (
	"net/http"

	_ "api/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

const swaggerPath = "/swagger/"

type SwaggerHandler struct {
	ui http.Handler
}

func NewSwaggerHandler() *SwaggerHandler {
	return &SwaggerHandler{
		ui: httpSwagger.Handler(
			httpSwagger.URL("doc.json"),
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("list"),
			httpSwagger.DefaultModelsExpandDepth(httpSwagger.HideModel),
		),
	}
}

func (h *SwaggerHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /swagger", h.Redirect)
	mux.HandleFunc("GET "+swaggerPath, h.UI)
}

func (h *SwaggerHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, swaggerPath, http.StatusMovedPermanently)
}

func (h *SwaggerHandler) UI(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != swaggerPath {
		h.ui.ServeHTTP(w, r)
		return
	}

	request := r.Clone(r.Context())
	request.URL.Path = swaggerPath + "index.html"
	request.RequestURI = request.URL.RequestURI()
	h.ui.ServeHTTP(w, request)
}
