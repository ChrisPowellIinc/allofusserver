package migrations

import (
	"github.com/ChrisPowellIinc/allofusserver/internal/config"
	"github.com/ChrisPowellIinc/allofusserver/models"
)

// MakeMigrations will create and modify existing tables without deleting existing ones
func MakeMigrations(con *config.Config) {
	// con := config.GetConf(false, false)
	db := con.DB
	db.AutoMigrate(&models.User{})
}
