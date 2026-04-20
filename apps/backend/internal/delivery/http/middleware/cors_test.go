package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORSAllowedOriginSetsHeaders(t *testing.T) {
	cfg := CORSConfig{
		AllowedOrigins: map[string]struct{}{
			"http://localhost:3111": {},
		},
		AllowedMethods:   "GET, POST, OPTIONS",
		AllowedHeaders:   "Accept, Authorization, Content-Type",
		AllowCredentials: true,
		PreflightStatus:  http.StatusNoContent,
	}

	handler := NewCORS(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://localhost:3111")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3111" {
		t.Fatalf("expected Access-Control-Allow-Origin to echo allowed origin, got %q", got)
	}

	if got := rec.Header().Get("Access-Control-Allow-Credentials"); got != "true" {
		t.Fatalf("expected Access-Control-Allow-Credentials true, got %q", got)
	}
}

func TestCORSDisallowedOriginOmitsHeaders(t *testing.T) {
	cfg := CORSConfig{
		AllowedOrigins: map[string]struct{}{
			"http://localhost:3111": {},
		},
		AllowedMethods:  "GET, POST, OPTIONS",
		AllowedHeaders:  "Accept, Authorization, Content-Type",
		PreflightStatus: http.StatusNoContent,
	}

	handler := NewCORS(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Origin", "http://malicious.local")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("expected Access-Control-Allow-Origin to be omitted, got %q", got)
	}
}

func TestCORSOptionsPreflightForAllowedOrigin(t *testing.T) {
	cfg := CORSConfig{
		AllowedOrigins: map[string]struct{}{
			"http://localhost:3111": {},
		},
		AllowedMethods:   "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowedHeaders:   "Accept, Authorization, Content-Type, X-Requested-With",
		AllowCredentials: true,
		PreflightStatus:  http.StatusNoContent,
	}

	handler := NewCORS(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Fatal("preflight request should not reach next handler")
	}))

	req := httptest.NewRequest(http.MethodOptions, "/", nil)
	req.Header.Set("Origin", "http://localhost:3111")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}

	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "http://localhost:3111" {
		t.Fatalf("expected Access-Control-Allow-Origin to echo allowed origin, got %q", got)
	}

	if got := rec.Header().Get("Access-Control-Allow-Methods"); got != cfg.AllowedMethods {
		t.Fatalf("expected Access-Control-Allow-Methods %q, got %q", cfg.AllowedMethods, got)
	}

	if got := rec.Header().Get("Access-Control-Allow-Headers"); got != cfg.AllowedHeaders {
		t.Fatalf("expected Access-Control-Allow-Headers %q, got %q", cfg.AllowedHeaders, got)
	}
}
