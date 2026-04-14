package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func getenv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func getenvInt(key string, def int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer: %w", key, err)
	}
	return n, nil
}

func main() {
	host := getenv("APP_HOST", "0.0.0.0")
	port, err := getenvInt("APP_PORT", 8080)
	if err != nil {
		log.Fatalf("config error : %v", err)
	}
	readHeaderTimeoutSec, err := getenvInt("APP_READ_HEADER_TIMEOUT_SEC", 5)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong\n"))
	})

	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: time.Duration(readHeaderTimeoutSec) * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("listening on http://%s/ping", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("received signal: %s", sig.String())
	case err := <-errCh:
		log.Printf("server error: %v", err)
	}

	shutdownTimeoutSec, err := getenvInt("APP_SHUTDOWN_TIMEOUT_SEC", 10)
	if err != nil {
		log.Fatalf("config error: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(shutdownTimeoutSec)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}

