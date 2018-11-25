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
	// lists all users
	router.Get("/", handler.Get)
	// Create a new user
	router.Post("/", handler.Get)
	// Update a new user
	router.Put("/", handler.Get)
	// Delete a user
	router.Delete("/", handler.Get)

	return router
}
