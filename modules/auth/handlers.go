package auth

import (
	"encoding/json"
	"log"
	"net/http"

	jwt "github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	"github.com/ChrisPowellIinc/allofusserver/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
)

// Login : Logs in an existing user
func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	res := models.Response{}
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}
	// Check if the user with that username exists
	db := handler.config.DB
	userPassword := user.Password
	// Get first matched record
	dd := db.Where("username = ?", user.Username).First(&user)
	if dd.Error != nil {
		log.Println(dd.Error)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPassword))
	if err != nil {
		log.Println("Passwords do not match")
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}

	_, token, err := jwt.TokenAuth.Encode(jwtauth.Claims{"user_id": user.ID})

	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusInternalServerError
		render.JSON(w, r, res)
		return
	}

	res = models.Response{
		Message: "Login Successful",
		Status:  http.StatusOK,
		Data:    map[string]string{"token": token},
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// Register : Registers a user
func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}
	errs := user.Validate(handler.config)
	if len(errs) > 0 {
		log.Println(errs)
		res.Message = "There are some problems in your forms"
		res.Status = http.StatusBadRequest
		res.Data = errs
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		return
	}

	db := handler.config.DB
	db.Create(&user)
	isNew := db.NewRecord(user) // return `false` after `user` created
	if isNew {
		res.Message = "There are some problems creating this user, please try again"
		res.Status = http.StatusInternalServerError
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		return
	}

	res.Message = "User created successfully"
	res.Status = http.StatusCreated
	// Add data if you need the user to sign in immediately.
	// res.Data =
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}
