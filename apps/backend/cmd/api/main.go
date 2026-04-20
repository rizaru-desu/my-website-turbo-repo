package main

import (
	"api/config"
	"api/internal/delivery/http/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Go Clean Architecture Berjalan!")
	})

	securityConfig := config.DefaultSecurityConfig()
	corsConfig := config.DefaultCORSConfig()

	// Apply CORS before security headers so preflight requests are answered cleanly.
	handler := middleware.NewCORS(corsConfig)(middleware.NewSecurityHeaders(securityConfig)(mux))

	fmt.Println("Server berjalan di http://localhost:3333")
	if err := http.ListenAndServe(":3333", handler); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
