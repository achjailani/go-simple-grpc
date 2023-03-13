package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunHTTPServer is a function to run http server
func RunHTTPServer(h http.Handler, addr string, shutdownTimeout time.Duration) error {
	log.Println("HTTP Server is starting ...")

	server := &http.Server{
		Handler:      h,
		Addr:         fmt.Sprintf(":%s", addr),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	server.SetKeepAlivesEnabled(false)
	serverError := make(chan error, 1)

	go func() {
		log.Printf("HTTP server is running at %v\n", server.Addr)
		serverError <- server.ListenAndServe()
	}()

	shutdownListenerChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownListenerChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		if err != nil {
			log.Fatalf("cannot start HTTP server, %v", err)
		}
	case sig := <-shutdownListenerChannel:
		log.Printf("HTTP server shutdown by signal: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Println("HTTP server shutdown by signal")

			err = server.Close()
			return err
		}

		log.Println("HTTP server shutdown by signal")
	}

	return nil
}
