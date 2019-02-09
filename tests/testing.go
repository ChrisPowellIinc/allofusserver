package tests

import (
	"log"

	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"

	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	jwtgo "github.com/dgrijalva/jwt-go"
)

var Con *config.Config

func init() {
	Con = GetConfig()
}

// GetConfig returns a testing config to be used by every tests
func GetConfig() *config.Config {
	return config.GetConf(true, false)
}

func GetAuth() string {
	jwt.Register([]byte(Con.Constants.JWTSecret))
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, jwtgo.MapClaims{"user_email": "odohi.david@gmail.com"})
	tokenString, err := token.SignedString([]byte(Con.Constants.JWTSecret))
	if err != nil {
		log.Println(err)
		return ""
	}
	return tokenString
}
