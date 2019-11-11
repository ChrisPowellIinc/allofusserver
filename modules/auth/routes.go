package auth

import (
	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
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
	authGroup := router.Group(nil)
	authGroup.Use(jwtauth.Verifier(jwt.TokenAuth))
	authGroup.Use(jwt.AuthHandler)
	// authGroup.Get("/", handler.GetLoggedInUser)
	// new user registrations route
	router.Post("/register", handler.Register)
	// login user route
	router.Post("/login", handler.Login)
	authGroup.Get("/sessionid", handler.GetStripeSessionID)
	authGroup.Get("/cards", handler.MyCards)
	authGroup.Delete("/cards/{pmID}", handler.DeleteCard)
	router.Post("/stripe/webhook/endpoint", handler.StripeWebhookHandler)

	return router
}
