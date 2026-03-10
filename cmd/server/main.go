package main

import (
	"log"
	"net/http"
	"os"

	"github.com/portal-interview/grill-master-service/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("POST /orders", handler.CreateOrder)
	mux.HandleFunc("GET /orders/{id}", handler.GetOrder)
	mux.HandleFunc("GET /stations", handler.ListStations)

	log.Printf("Grill Master Service starting on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}
