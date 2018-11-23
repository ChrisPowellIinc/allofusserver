package user

import (
	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/go-chi/chi"
)

// Handler : Routes handler
type Handler struct {
	config *config.Config
}

var handler *Handler

// New : Creates a new handler object
func New(config *config.Config) *Handler {
	return &Handler{config: config}
}

// Routes : Defines API routes for this module
func Routes(config *config.Config) *chi.Mux {
	handler = New(config)
	router := chi.NewRouter()
	router.Get("/get", handler.Get)

	return router
}
