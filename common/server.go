package common

import (
	"net/http"
	"time"
)

const (
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	MaxHeaderBytes = 1 << 20
)

func NewServer(handler http.Handler, port string) *http.Server {
	return &http.Server{
		Addr:           port,
		Handler:        handler,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}
}