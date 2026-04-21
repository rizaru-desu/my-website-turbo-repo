package config

import (
	"net/http"
	"strings"

	"api/internal/delivery/http/middleware"
)

const contentSecurityPolicy = "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' data:; connect-src 'self'; frame-ancestors 'self'; base-uri 'self'; form-action 'self'"
const developmentContentSecurityPolicy = "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' data:; connect-src 'self'; frame-ancestors 'self'; base-uri 'self'; form-action 'self'"

func DefaultSecurityConfig() middleware.SecurityConfig {
	return middleware.SecurityConfig{
		ContentSecurityPolicy:         contentSecurityPolicy,
		FrameOptions:                  "SAMEORIGIN",
		ContentTypeOptions:            "nosniff",
		ReferrerPolicy:                "strict-origin-when-cross-origin",
		PermissionsPolicy:             "accelerometer=(), camera=(), geolocation=(), gyroscope=(), magnetometer=(), microphone=(), payment=(), usb=()",
		CrossOriginOpenerPolicy:       "same-origin",
		CrossOriginResourcePolicy:     "same-origin",
		PermittedCrossDomainPolicies:  "none",
		EnableStrictTransportSecurity: true,
		StrictTransportSecurityMaxAge: "max-age=31536000; includeSubDomains",
	}
}

func SecurityConfigForEnvironment(environment string) middleware.SecurityConfig {
	config := DefaultSecurityConfig()
	if strings.EqualFold(strings.TrimSpace(environment), "development") {
		config.ContentSecurityPolicy = developmentContentSecurityPolicy
	}

	return config
}

func DefaultCORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowedOrigins: map[string]struct{}{
			"http://127.0.0.1:3111": {},
			"http://127.0.0.1:3222": {},
			"http://localhost:3111": {},
			"http://localhost:3222": {},
		},
		AllowedMethods:   "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowedHeaders:   "Accept, Authorization, Content-Type, X-Requested-With",
		AllowCredentials: true,
		PreflightStatus:  http.StatusNoContent,
	}
}
