package main

import (
	"log"
	"net/http"

	"github.com/IvanSR01/infrapilot/internal/server"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /heathz", server.HealthHandler)

	log.Println("Server started on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
