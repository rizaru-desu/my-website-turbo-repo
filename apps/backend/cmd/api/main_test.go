package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootRouteReturnsOK(t *testing.T) {
	handler := newHandler()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestV1HealthRouteReturnsOK(t *testing.T) {
	handler := newHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestUnknownRouteReturnsNotFound(t *testing.T) {
	handler := newHandler()
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestSwaggerRoutesAreAvailableInDevelopment(t *testing.T) {
	t.Setenv("APP_ENV", "development")

	handler := newHandler()

	uiReq := httptest.NewRequest(http.MethodGet, "/swagger/", nil)
	uiRec := httptest.NewRecorder()
	handler.ServeHTTP(uiRec, uiReq)
	if uiRec.Code != http.StatusOK {
		t.Fatalf("expected swagger UI status %d, got %d", http.StatusOK, uiRec.Code)
	}

	docReq := httptest.NewRequest(http.MethodGet, "/swagger/doc.json", nil)
	docRec := httptest.NewRecorder()
	handler.ServeHTTP(docRec, docReq)
	if docRec.Code != http.StatusOK {
		t.Fatalf("expected swagger spec status %d, got %d", http.StatusOK, docRec.Code)
	}

	var spec map[string]any
	if err := json.Unmarshal(docRec.Body.Bytes(), &spec); err != nil {
		t.Fatalf("expected swagger spec to be valid JSON, got error: %v", err)
	}

	if spec["swagger"] != "2.0" {
		t.Fatalf("expected Swagger 2.0, got %v", spec["swagger"])
	}
}

func TestSwaggerRoutesAreUnavailableOutsideDevelopment(t *testing.T) {
	t.Setenv("APP_ENV", "production")

	handler := newHandler()

	for _, path := range []string{"/swagger/", "/swagger/doc.json"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Fatalf("expected %s to return status %d, got %d", path, http.StatusNotFound, rec.Code)
		}
	}
}
