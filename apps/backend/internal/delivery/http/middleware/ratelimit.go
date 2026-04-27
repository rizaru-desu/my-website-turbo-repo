package middleware

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type IPRateLimitConfig struct {
	Window time.Duration
	Max    int
	Routes map[string]struct{}
	Now    func() time.Time
}

type ipRateLimitEntry struct {
	windowStart time.Time
	count       int
}

func NewIPRateLimiter(cfg IPRateLimitConfig) func(http.Handler) http.Handler {
	if cfg.Window <= 0 {
		cfg.Window = time.Minute
	}
	if cfg.Max <= 0 {
		cfg.Max = 5
	}
	if cfg.Now == nil {
		cfg.Now = time.Now
	}

	var mu sync.Mutex
	attempts := map[string]ipRateLimitEntry{}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !rateLimitApplies(cfg.Routes, r) {
				next.ServeHTTP(w, r)
				return
			}

			now := cfg.Now().UTC()
			key := r.Method + " " + r.URL.Path + " " + clientIP(r)

			mu.Lock()
			entry := attempts[key]
			if entry.windowStart.IsZero() || !entry.windowStart.Add(cfg.Window).After(now) {
				entry = ipRateLimitEntry{windowStart: now}
			}
			entry.count++
			allowed := entry.count <= cfg.Max
			attempts[key] = entry
			cleanupExpiredRateLimitEntries(attempts, now, cfg.Window)
			mu.Unlock()

			if !allowed {
				http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func rateLimitApplies(routes map[string]struct{}, r *http.Request) bool {
	if len(routes) == 0 {
		return true
	}

	_, ok := routes[r.Method+" "+r.URL.Path]
	return ok
}

func cleanupExpiredRateLimitEntries(entries map[string]ipRateLimitEntry, now time.Time, window time.Duration) {
	for key, entry := range entries {
		if entry.windowStart.Add(window).Before(now) {
			delete(entries, key)
		}
	}
}

func clientIP(r *http.Request) string {
	if forwarded := strings.TrimSpace(r.Header.Get("X-Forwarded-For")); forwarded != "" {
		ip, _, _ := strings.Cut(forwarded, ",")
		return strings.TrimSpace(ip)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return host
	}

	return r.RemoteAddr
}
