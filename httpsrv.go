package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func httpsrvInit(addr string, logger *log.Logger,
	handler func(http.ResponseWriter, *http.Request)) (*http.Server, chan error) {
	listenErrChan := make(chan error)
	server := &http.Server{
		Addr:         addr,
		Handler:      http.HandlerFunc(handler),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	go func() {
		listenErrChan <- server.ListenAndServe()
	}()
	return server, listenErrChan
}

func httpsrvClose(server *http.Server, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	server.SetKeepAlivesEnabled(false)
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("Could not gracefully shutdown the http server: %v", err)
	}
	return nil
}
