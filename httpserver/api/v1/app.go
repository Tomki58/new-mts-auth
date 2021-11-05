package v1

import (
	"mts/auth/config"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// App is a struct for representing business-logic.
type App struct {
	Config *config.Config
}

// New creates and returns a new object of App.
// It contains credentials
func New() (*App, error) {
	config, err := config.New("./credentials/creds.txt")
	if err != nil {
		return nil, err
	}
	return &App{
		Config: config,
	}, nil
}

// ApplyEndpoints sets routes for router.
func (a *App) ApplyEndpoints(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.Logger, middleware.BasicAuth("User client realm", a.Config.Credentials))
		// TODO: swap handler to login handler
		r.Get("/login", func(rw http.ResponseWriter, r *http.Request) {
			rw.Write([]byte("Welcome string!"))
		})
	})
}
