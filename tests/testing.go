package tests

import (
	"github.com/ChrisPowellIinc/allofusserver/internal/config"
)

var Con *config.Config

func init() {
	Con = GetConfig()
}

// GetConfig returns a testing config to be used by every tests
func GetConfig() *config.Config {
	return config.GetConf(true, false)
}
