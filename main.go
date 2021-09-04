package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/FollowTheProcess/go_microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api ", log.LstdFlags)
	productsHandler := handlers.NewProducts(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", productsHandler)

	// Can define a custom server with specific settings
	server := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  time.Second * 20,
		ReadTimeout:  time.Second * 1,
		WriteTimeout: time.Second * 1,
	}

	// Starts the server in the background
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Trap sigterm or sigkill and gracefully shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// Block until signal is received
	sig := <-sigChan
	logger.Println("Got signal: ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Println("Error shutting down server: ", err)
		os.Exit(1)
	}
}
