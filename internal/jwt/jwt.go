package jwt

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// TokenAuth : JWTAuth object
var TokenAuth *jwtauth.JWTAuth

func Register(secret []byte) {
	TokenAuth = jwtauth.New("HS256", secret, nil)
}

// AuthHandler : Custom handler, handles authentication requests returns JSON
func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if token == nil || !token.Valid || err != nil {
			err := make(map[string]string)
			err["error"] = "Unauthorized"

			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GetLoggedInUserEmail get the logged in user email which is a string
func GetLoggedInUserEmail(context context.Context) (string, error) {
	_, claims, err := jwtauth.FromContext(context)
	if err != nil {
		return "", err
	}
	email, ok := claims["user_email"].(string)
	if !ok {
		return "", errors.New("Invalid User ID gotten from claims")
	}
	return email, nil
}
