package v1

import (
	"github.com/go-chi/chi"
)

// App is a struct for representing business-logic
type App struct{}

// New creates and returns a new object of App
func New() (*App, error) {
	return &App{}, nil
}

// ApplyEndpoints sets routes for router
func (a *App) ApplyEndpoints(router *chi.Mux) {
}
