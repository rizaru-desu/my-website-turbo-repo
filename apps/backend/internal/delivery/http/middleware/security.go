package middleware

import (
	"net/http"
	"strings"
)

type SecurityConfig struct {
	ContentSecurityPolicy         string
	FrameOptions                  string
	ContentTypeOptions            string
	ReferrerPolicy                string
	PermissionsPolicy             string
	CrossOriginOpenerPolicy       string
	CrossOriginResourcePolicy     string
	PermittedCrossDomainPolicies  string
	EnableStrictTransportSecurity bool
	StrictTransportSecurityMaxAge string
}

func NewSecurityHeaders(cfg SecurityConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headers := w.Header()

			setHeader(headers, "Content-Security-Policy", cfg.ContentSecurityPolicy)
			setHeader(headers, "X-Frame-Options", cfg.FrameOptions)
			setHeader(headers, "X-Content-Type-Options", cfg.ContentTypeOptions)
			setHeader(headers, "Referrer-Policy", cfg.ReferrerPolicy)
			setHeader(headers, "Permissions-Policy", cfg.PermissionsPolicy)
			setHeader(headers, "Cross-Origin-Opener-Policy", cfg.CrossOriginOpenerPolicy)
			setHeader(headers, "Cross-Origin-Resource-Policy", cfg.CrossOriginResourcePolicy)
			setHeader(headers, "X-Permitted-Cross-Domain-Policies", cfg.PermittedCrossDomainPolicies)

			if cfg.EnableStrictTransportSecurity && isHTTPS(r) {
				setHeader(headers, "Strict-Transport-Security", cfg.StrictTransportSecurityMaxAge)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func setHeader(headers http.Header, key, value string) {
	if value == "" {
		return
	}

	headers.Set(key, value)
}

func isHTTPS(r *http.Request) bool {
	if r.TLS != nil {
		return true
	}

	if strings.EqualFold(r.Header.Get("X-Forwarded-Proto"), "https") {
		return true
	}

	return strings.Contains(strings.ToLower(r.Header.Get("Forwarded")), "proto=https")
}
