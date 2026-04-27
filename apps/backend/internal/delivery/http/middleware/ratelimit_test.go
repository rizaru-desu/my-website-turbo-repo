package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestIPRateLimiterLimitsConfiguredRoute(t *testing.T) {
	now := time.Date(2026, 4, 27, 8, 0, 0, 0, time.UTC)
	handler := NewIPRateLimiter(IPRateLimitConfig{
		Window: time.Minute,
		Max:    2,
		Routes: map[string]struct{}{
			"POST /api/v1/auth/email-verification": {},
		},
		Now: func() time.Time { return now },
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/email-verification", nil)
		req.RemoteAddr = "192.0.2.1:1234"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected request %d status %d, got %d", i+1, http.StatusOK, rec.Code)
		}
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/email-verification", nil)
	req.RemoteAddr = "192.0.2.1:1234"
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusTooManyRequests {
		t.Fatalf("expected status %d after limit, got %d", http.StatusTooManyRequests, rec.Code)
	}
}

func TestIPRateLimiterResetsAfterWindow(t *testing.T) {
	now := time.Date(2026, 4, 27, 8, 0, 0, 0, time.UTC)
	handler := NewIPRateLimiter(IPRateLimitConfig{
		Window: time.Minute,
		Max:    1,
		Routes: map[string]struct{}{
			"POST /api/v1/auth/forgot-password": {},
		},
		Now: func() time.Time { return now },
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/forgot-password", nil)
	req.RemoteAddr = "192.0.2.1:1234"
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected first request status %d, got %d", http.StatusOK, rec.Code)
	}

	now = now.Add(time.Minute + time.Second)

	req = httptest.NewRequest(http.MethodPost, "/api/v1/auth/forgot-password", nil)
	req.RemoteAddr = "192.0.2.1:1234"
	rec = httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected request after window status %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestIPRateLimiterSkipsUnconfiguredRoutes(t *testing.T) {
	handler := NewIPRateLimiter(IPRateLimitConfig{
		Window: time.Minute,
		Max:    1,
		Routes: map[string]struct{}{
			"POST /api/v1/auth/email-verification": {},
		},
	})(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for i := 0; i < 3; i++ {
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
		req.RemoteAddr = "192.0.2.1:1234"
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusOK {
			t.Fatalf("expected unconfigured request %d status %d, got %d", i+1, http.StatusOK, rec.Code)
		}
	}
}
