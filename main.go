package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	port := "8080"

	// 404 Handler
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})

	// Server setup
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	//start the server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()
	log.Println("Server is ready to handle requests at :8080")

	<-stop
	log.Println("Server is shutting down...")

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
