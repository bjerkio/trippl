package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	_, b, _, _      = runtime.Caller(0)
	basePath        = filepath.Dir(b)
	projectBasePath = filepath.Join(basePath, "..", "..", "..", ".env")
)

// GetConfig returns config from environment variables
func GetConfig() (*TripplConfig, error) {
	err := godotenv.Load(projectBasePath)
	if err != nil {
		log.Println("Error loading .env file")
	}

	var c TripplConfig
	err = envconfig.Process("trippl", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
