package models

import (
	"strings"

	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/jinzhu/gorm"
)

// User : User model
type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:varchar(100)"`
	LastName  string `json:"last_name" gorm:"type:varchar(100)"`
	Email     string `json:"email" gorm:"type:varchar(100);unique_index"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

// FindByID : Returns user with the matching ID
func (User) FindByID(c *config.Config, id uint) (User, error) {
	db := c.DB
	user := User{}

	if err := db.Where(map[string]interface{}{"ID": id}).
		First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

// FindByEmail queries the DB for a user with the specified email address
func (u User) FindByEmail(c *config.Config) (User, error) {
	db := c.DB
	user := User{}
	err := db.Where("email = ?", u.Email).First(&user).Error
	return user, err
}

// FindByUsername queries the DB for a user with the specified username
func (u User) FindByUsername(c *config.Config) (User, error) {
	db := c.DB
	user := User{}
	err := db.Where("username = ?", u.Username).First(&user).Error
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

	if len(u.Username) <= 0 {
		data["username"] = "You must enter a valid username"
	} else {
		// check if username exists...
		_, err := u.FindByUsername(c)
		if err == nil {
			data["username"] = "Username is taken"
		}
	}

	if len(u.Password) <= 6 {
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
