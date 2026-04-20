package middleware

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecurityHeadersOnNormalResponse(t *testing.T) {
	cfg := SecurityConfig{
		ContentSecurityPolicy:         "default-src 'self'",
		FrameOptions:                  "SAMEORIGIN",
		ContentTypeOptions:            "nosniff",
		ReferrerPolicy:                "strict-origin-when-cross-origin",
		PermissionsPolicy:             "camera=()",
		CrossOriginOpenerPolicy:       "same-origin",
		CrossOriginResourcePolicy:     "same-origin",
		PermittedCrossDomainPolicies:  "none",
		EnableStrictTransportSecurity: true,
		StrictTransportSecurityMaxAge: "max-age=31536000; includeSubDomains",
	}

	handler := NewSecurityHeaders(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if got := rec.Header().Get("Content-Security-Policy"); got != cfg.ContentSecurityPolicy {
		t.Fatalf("expected Content-Security-Policy %q, got %q", cfg.ContentSecurityPolicy, got)
	}

	if got := rec.Header().Get("X-Frame-Options"); got != cfg.FrameOptions {
		t.Fatalf("expected X-Frame-Options %q, got %q", cfg.FrameOptions, got)
	}

	if got := rec.Header().Get("X-Content-Type-Options"); got != cfg.ContentTypeOptions {
		t.Fatalf("expected X-Content-Type-Options %q, got %q", cfg.ContentTypeOptions, got)
	}

	if got := rec.Header().Get("Referrer-Policy"); got != cfg.ReferrerPolicy {
		t.Fatalf("expected Referrer-Policy %q, got %q", cfg.ReferrerPolicy, got)
	}

	if got := rec.Header().Get("Strict-Transport-Security"); got != "" {
		t.Fatalf("expected Strict-Transport-Security to be omitted for plain HTTP, got %q", got)
	}
}

func TestSecurityHeadersOnlySetHSTSForHTTPS(t *testing.T) {
	cfg := SecurityConfig{
		EnableStrictTransportSecurity: true,
		StrictTransportSecurityMaxAge: "max-age=31536000; includeSubDomains",
	}

	testCases := []struct {
		name            string
		request         *http.Request
		expectedPresent bool
	}{
		{
			name:            "direct tls request",
			request:         httptest.NewRequest(http.MethodGet, "/", nil),
			expectedPresent: true,
		},
		{
			name:            "proxied https request",
			request:         httptest.NewRequest(http.MethodGet, "/", nil),
			expectedPresent: true,
		},
		{
			name:            "plain http request",
			request:         httptest.NewRequest(http.MethodGet, "/", nil),
			expectedPresent: false,
		},
	}

	testCases[0].request.TLS = &tls.ConnectionState{}
	testCases[1].request.Header.Set("X-Forwarded-Proto", "https")

	handler := NewSecurityHeaders(cfg)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, tc.request)

			got := rec.Header().Get("Strict-Transport-Security")
			if tc.expectedPresent && got != cfg.StrictTransportSecurityMaxAge {
				t.Fatalf("expected Strict-Transport-Security %q, got %q", cfg.StrictTransportSecurityMaxAge, got)
			}

			if !tc.expectedPresent && got != "" {
				t.Fatalf("expected Strict-Transport-Security to be omitted, got %q", got)
			}
		})
	}
}
