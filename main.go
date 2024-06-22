package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"exoplanet_microservice/handlers"
	"exoplanet_microservice/repository"
	"exoplanet_microservice/services"

	"github.com/gorilla/mux"
)

func main() {

	repo := repository.NewInMemoryExoplanetRepository()
	service := services.NewExoplanetService(repo)
	handler := handlers.NewExoplanetHandler(service)

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/exoplanets", handler.AddExoplanet).Methods("POST")
	r.HandleFunc("/exoplanets", handler.ListExoplanets).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handler.GetExoplanetByID).Methods("GET")
	r.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanet).Methods("PUT")
	r.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanet).Methods("DELETE")
	r.HandleFunc("/exoplanets/{id}/fuel", handler.EstimateFuel).Methods("GET")

	// 404 Handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 not found", http.StatusNotFound)
	})

	// Server setup
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Goroutine to start the server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	log.Println("Server is ready to handle requests at :8080")

	// Blocking until we receive a signal
	<-stop
	log.Println("Server is shutting down...")

	// Attempt to gracefully shutdown the server
	if err := gracefulShutdown(srv, 5*time.Second); err != nil {
		log.Fatalf("Could not gracefully shut down the server: %v\n", err)
	}
	log.Println("Server stopped")
}

func gracefulShutdown(server *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return server.Shutdown(ctx)
}
