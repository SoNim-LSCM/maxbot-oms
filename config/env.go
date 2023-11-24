package config

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const projectDirName = "maxbot_oms" // change to relevant project name

// this function will load the .env file if the GO_ENV environment variable is not set
func LoadENV() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	workingPath := filepath.Dir(ex)
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "development" {
		err = godotenv.Load("maxbot_oms.env")
	} else {
		err = godotenv.Load(workingPath + `/maxbot_oms.env`)
	}
	if err != nil {
		return err
	}

	return nil
}
