package main

import (
	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	logger "github.com/ChrisPowellIinc/allofusserver/internal/log"
	"github.com/ChrisPowellIinc/allofusserver/models"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	con := config.GetConf(false, false)
	db := con.DB
	user := models.User{
		FirstName:      "Edem",
		LastName:       "Akpan",
		Username:       "edem",
		Email:          "edem@gmail.com",
		PasswordString: "password",
	}
	// Create
	var err error
	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Fatal("Could not generate password")
	}
	db.C("user").Insert(&user)
}
