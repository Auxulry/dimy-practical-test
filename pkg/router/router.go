// Package router describe all configurations
package router

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(CORS)
	router.Use(Recovery)

	return router
}