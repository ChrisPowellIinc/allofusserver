package user

import (
	"net/http"

	"github.com/go-chi/render"
)

// Get : Shows that the app is working
func (handler *Handler) Get(w http.ResponseWriter, r *http.Request) {

	// return a user's timeline feed (POSTS)
	type resStruct struct {
		Message string `json:"msg"`
	}

	res := resStruct{
		Message: "It works!",
	}

	render.JSON(w, r, res)

	return
}
