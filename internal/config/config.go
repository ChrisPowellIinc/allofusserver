package config

import (
	"encoding/json"
	"log"
	"path/filepath"
	"runtime"

	// Imports mysql driver for the effect
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config : Config type
type Config struct {
	Constants
	DB *gorm.DB
}

// Constants type holds data defined within the viper config file
type Constants struct {
	DBURL     string
	PORT      string
	ENV       string
	LogDir    string
	JWTSecret string
}

func parseConfigFile(isTesting, debug bool) Constants {
	_, filename, _, _ := runtime.Caller(0)
	viper.AddConfigPath(filepath.Join(filename + "/../../../"))

	if isTesting {
		viper.SetConfigName("app.test.config")
	} else {
		viper.SetConfigName("app.config")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Print(err)
	}

	constants := Constants{}
	err = viper.Unmarshal(&constants)
	if err != nil {
		log.Panicf("Unable to parse config file, %v", err)
	}

	if debug {
		log.Println("In Debug Mode. Logging configuration data")
		indentedConfig, _ := json.MarshalIndent(constants, "", "\t")
		log.Printf("\n %s \n \n", indentedConfig)
	}

	return constants
}

// GetConf : Returns a pointer to the config object
func GetConf(isTest, isDebug bool) *Config {
	var config Config
	config.Constants = parseConfigFile(isTest, isDebug)

	var err error
	config.DB, err = gorm.Open("mysql", config.Constants.DBURL)

	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to database"))
	}

	config.DB.SingularTable(true)
	return &config
}
