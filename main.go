package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Gprisco/decanto-golang/env"
	"github.com/Gprisco/decanto-golang/routers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "main", log.LstdFlags)

	r := mux.NewRouter()

	routers.Config(r)

	config := env.GetInstance()

	s := &http.Server{
		Addr:         config.Port,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Put The Server start in a goroutine, so it will not block
	go func() {
		log.Println("Server listening...")
		err := s.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	// Create a channel
	sigChan := make(chan os.Signal, 1) // 1 -> buffer of size 1

	// Notify sigChan everytime we get Interrupt or Kill Signal
	signal.Notify(sigChan, os.Interrupt)

	// !!! THIS IS BLOCKING -> will listen for signals (specified above by us)
	sig := <-sigChan
	logger.Println("Gracefully shutdown...", sig)

	// Get Background context, assign a 30 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server with the timeout specified above
	s.Shutdown(ctx)
}
