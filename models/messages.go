package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type Response struct {
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
	Data    map[string]interface{} `json:"data"`
}

func HandleResponse(w http.ResponseWriter, r *http.Request, message string, status int) {
	res := Response{}
	res.Message = message
	res.Status = status
	render.Status(r, status)
	render.JSON(w, r, res)
}
