package server

import (
	"net/http"
	"time"
)

func NewServer(addr string) (*http.Server, *http.ServeMux) {

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server, mux
}
