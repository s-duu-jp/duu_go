package env

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	Env     = "development"
	config  map[string]string
	once    sync.Once
	loadErr error
)

// GetConfig returns the loaded configuration as a singleton
func GetConfig() (map[string]string, error) {
	once.Do(loadConfig)
	return config, loadErr
}

// loadConfig loads environment variables from the appropriate .env file based on the environment
func loadConfig() {
	var envFile string
	switch Env {
	case "production":
		envFile = "./config/env/conf/prod.env"
	case "staging":
		envFile = "./config/env/conf/stg.env"
	default:
		envFile = "./config/env/conf/dev.env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Error loading %s file", envFile)
		loadErr = err
		return
	}

	config = make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			config[pair[0]] = pair[1]
		}
	}

	log.Printf("Loaded %s file", envFile)
}
