package handler

import (
	systemusecase "api/internal/usecase/system"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootReturnsMinimalStatus(t *testing.T) {
	systemHandler := NewSystemHandler("1.0.0", "development", stubHealthReporter{})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	systemHandler.Root(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if got := rec.Header().Get("Content-Type"); got != "application/json; charset=utf-8" {
		t.Fatalf("expected JSON content type, got %q", got)
	}

	var response map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("expected valid JSON response, got error: %v", err)
	}

	if response["status"] != "ok" {
		t.Fatalf("expected generic ok status, got %v", response["status"])
	}

	for _, field := range []string{"name", "version", "environment", "description", "timestamp", "endpoints"} {
		if _, exists := response[field]; exists {
			t.Fatalf("expected root response not to expose %q", field)
		}
	}
}

func TestHealthReturnsOK(t *testing.T) {
	systemHandler := NewSystemHandler("1.0.0", "development", stubHealthReporter{})
	req := httptest.NewRequest(http.MethodGet, healthPath, nil)
	rec := httptest.NewRecorder()

	systemHandler.Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var response healthResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("expected valid JSON response, got error: %v", err)
	}

	if response.Version != "1.0.0" {
		t.Fatalf("expected version 1.0.0, got %q", response.Version)
	}

	if response.Status != "operational" {
		t.Fatalf("expected status operational, got %q", response.Status)
	}

	if response.SystemHealth.CPUUsage != "12.5%" {
		t.Fatalf("expected CPU usage to match stub, got %q", response.SystemHealth.CPUUsage)
	}

	if response.SystemHealth.Storage.Total != "14.9GB" {
		t.Fatalf("expected storage total to match stub, got %q", response.SystemHealth.Storage.Total)
	}
}

type stubHealthReporter struct{}

func (stubHealthReporter) Snapshot() systemusecase.HealthSnapshot {
	return systemusecase.HealthSnapshot{
		Status: "operational",
		SystemHealth: systemusecase.SystemHealth{
			CPUUsage: "12.5%",
			Memory: systemusecase.ResourceUsage{
				Used:       "64MB",
				Total:      "1024MB",
				Percentage: "6.2%",
			},
			Storage: systemusecase.StorageUsage{
				Used:       "2.1GB",
				Free:       "12.8GB",
				Total:      "14.9GB",
				Percentage: "14%",
			},
			Uptime: "5d 12h 30m",
		},
	}
}
