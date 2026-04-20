package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Go Clean Architecture Berjalan!")
	})

	fmt.Println("Server berjalan di http://localhost:3333")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		fmt.Println("Gagal menjalankan server:", err)
	}
}
