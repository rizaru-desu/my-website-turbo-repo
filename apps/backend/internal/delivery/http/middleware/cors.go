package middleware

import "net/http"

type CORSConfig struct {
	AllowedOrigins   map[string]struct{}
	AllowedMethods   string
	AllowedHeaders   string
	AllowCredentials bool
	PreflightStatus  int
}

func NewCORS(cfg CORSConfig) func(http.Handler) http.Handler {
	preflightStatus := cfg.PreflightStatus
	if preflightStatus == 0 {
		preflightStatus = http.StatusNoContent
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if _, ok := cfg.AllowedOrigins[origin]; ok {
				headers := w.Header()
				headers.Set("Access-Control-Allow-Origin", origin)
				setHeader(headers, "Access-Control-Allow-Methods", cfg.AllowedMethods)
				setHeader(headers, "Access-Control-Allow-Headers", cfg.AllowedHeaders)
				if cfg.AllowCredentials {
					headers.Set("Access-Control-Allow-Credentials", "true")
				}
				headers.Set("Vary", "Origin")
			}

			if r.Method == http.MethodOptions {
				w.WriteHeader(preflightStatus)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
