package main

import (
	"log"
	"net/http"
	"context"
	
	"github.com/IvanSR01/infrapilot/internal/server"	
	"github.com/IvanSR01/infrapilot/internal/storage"
)

func main() {
	ctx := context.Background()

	pool, err := storage.NewPostgres(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer pool.Close()

	log.Println("PostgreSQL connected")
	
	mux := http.NewServeMux()

	mux.HandleFunc("GET /heathz", server.HealthHandler)

	log.Println("Server started on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
