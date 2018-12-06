package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/globalsign/mgo/bson"

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
	userPassword := user.PasswordString
	// Get first matched record
	err = db.C("user").Find(bson.M{"username": user.Username}).One(&user)
	if err != nil {
		log.Println(err)
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

	_, token, err := jwt.TokenAuth.Encode(jwtauth.Claims{"user_email": user.Email})

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
		Data: map[string]interface{}{
			"token":      token,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"phone":      user.Phone,
			"email":      user.Email,
			"username":   user.Username,
			"image":      user.Image,
		},
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

	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.PasswordString), bcrypt.DefaultCost)
	if err != nil {
		res.Message = "Sorry a problem occured, please try again"
		res.Status = http.StatusInternalServerError
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		return
	}

	db := handler.config.DB
	err = db.C("user").Insert(&user)
	if err != nil {
		res.Message = "Sorry a problem occured, please try again"
		res.Status = http.StatusInternalServerError
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		return
	}

	res.Message = "User created successfully"
	res.Status = http.StatusCreated
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}
