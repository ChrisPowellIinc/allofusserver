package config

import (
	"encoding/json"
	"log"
	"path/filepath"
	"runtime"

	// Imports mysql driver for the effect
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Constants type holds data defined within the viper config file
type Constants struct {
	DB             string
	DBURL          string
	PORT           string
	ENV            string
	LogDir         string
	JWTSecret      string
	S3Region       string
	S3Bucket       string
	AWSAccessKeyID string
	AWSSecretKey   string
}

// Config : Config type
type Config struct {
	Constants
	DBSession  *mgo.Session
	DB         *mgo.Database
	AwsSession *session.Session
}

func parseConfigFile(isTesting, debug bool) Constants {
	viper.AddConfigPath(filepath.Join("./"))

	if isTesting {
		_, filename, _, _ := runtime.Caller(0)
		viper.AddConfigPath(filepath.Join(filename, "/../../../"))
		viper.SetConfigName("app.test.config")
	} else if debug {
		_, filename, _, _ := runtime.Caller(0)
		viper.AddConfigPath(filepath.Join(filename, "/../../../"))
		viper.SetConfigName("app.config")
	} else {
		viper.SetConfigName("app.prod.config")
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

	// if debug {
	log.Println("In Debug Mode. Logging configuration data")
	indentedConfig, _ := json.MarshalIndent(constants, "", "\t")
	log.Printf("\n %s \n \n", indentedConfig)
	// }

	return constants
}

// GetConf : Returns a pointer to the config object
func GetConf(isTest, isDebug bool) *Config {
	var config Config
	config.Constants = parseConfigFile(isTest, isDebug)

	var err error
	config.DBSession, err = mgo.Dial(config.Constants.DBURL)
	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to Mongo database"))
	}

	// create an AWS session which can be
	// reused if we're uploading many files
	s, err := session.NewSession(&aws.Config{
		Region: aws.String(config.Constants.S3Region),
		Credentials: credentials.NewStaticCredentials(
			config.Constants.AWSAccessKeyID, // id
			config.Constants.AWSSecretKey,   // secret
			""),                             // token can be left blank for now
	})
	if err != nil {
		log.Fatal(err)
	}
	config.AwsSession = s

	config.DB = config.DBSession.DB(config.Constants.DB)

	return &config
}
