package models

import (
	"strings"
	"time"

	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/globalsign/mgo/bson"
)

// User : User model
type User struct {
	FirstName      string    `json:"first_name" bson:"first_name,omitempty"`
	LastName       string    `json:"last_name" bson:"last_name,omitempty"`
	Phone          string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Email          string    `json:"email" bson:"email,omitempty"`
	Username       string    `json:"username" bson:"username,omitempty"`
	Password       []byte    `json:"-" bson:"password,omitempty"`
	PasswordString string    `json:"password,omitempty" bson:"-"`
	Reset          string    `json:"-" bson:"reset"`
	Image          string    `json:"image,omitempty" bson:"image,omitempty"`
	DateCreated    time.Time `json:"date_created,omitempty" bson:"date_created,omitempty"`
	AccessToken    string    `json:"token,omitempty" bson:"token,omitempty"`
}

// FindByPhone queries the DB for a user with the specified phone number
func (u User) FindByPhone(c *config.Config) (User, error) {
	user := User{}
	err := c.DB.C("user").Find(bson.M{"phone": u.Phone}).One(&user)
	return user, err
}

// FindByEmail queries the DB for a user with the specified email address
func (u User) FindByEmail(c *config.Config) (User, error) {
	user := User{}
	err := c.DB.C("user").Find(bson.M{"email": u.Email}).One(&user)
	return user, err
}

// FindByUsername queries the DB for a user with the specified username
func (u User) FindByUsername(c *config.Config) (User, error) {
	user := User{}
	err := c.DB.C("user").Find(bson.M{"username": u.Username}).One(&user)
	return user, err
}

// Validate returns a map of errors if the values provided are not valid
func (u User) Validate(c *config.Config) map[string]string {
	var data = map[string]string{}
	if len(u.FirstName) <= 0 {
		data["first_name"] = "You must enter a valid first name"
	}
	if len(u.LastName) <= 0 {
		data["last_name"] = "You must enter a valid last name"
	}

	if len(u.Phone) <= 10 {
		data["phone"] = "You must enter a valid phone number"
	} else {
		// check if phone number exists...
		_, err := u.FindByPhone(c)
		if err == nil {
			data["phone"] = "Phone number is taken"
		}
	}

	if len(u.Username) <= 0 {
		data["username"] = "You must enter a valid username"
	} else {
		// check if username exists...
		_, err := u.FindByUsername(c)
		if err == nil {
			data["username"] = "Username is taken"
		}
	}

	if len(u.PasswordString) <= 6 {
		data["password"] = "You must enter a valid password: minimum of 6 characters"
	}

	if len(u.Email) <= 0 || !strings.Contains(u.Email, "@") {
		data["email"] = "You must enter a valid email address"
	} else {
		// check if email exists...
		_, err := u.FindByEmail(c)
		if err == nil {
			data["email"] = "Email already exists"
		}
	}
	return data
}
