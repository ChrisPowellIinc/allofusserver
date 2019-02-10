package user

import (
	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	"github.com/go-chi/jwtauth"
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
	router.Use(jwtauth.Verifier(jwt.TokenAuth))
	router.Use(jwt.AuthHandler)
	// lists all users
	router.Get("/", handler.Get)
	// Create a new user
	router.Post("/", handler.Get)
	// Update a new user
	router.Put("/", handler.Get)
	// Delete a user
	router.Delete("/", handler.Get)
	// Upload image
	router.Post("/upload", handler.UploadProfilePic)

	return router
}
