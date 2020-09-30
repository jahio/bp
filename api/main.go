package main

import (
	"github.com/gobuffalo/pop"
	"jahio/bp/models"
	"log"
	"fmt"
	"os"
	"strconv"
)

type appConfig struct {
	Port    int
	Runtime string
}

func (a *appConfig) validRuntime(r string) bool {
	switch r {
	case
		"development", "test", "production":
		return true
	}
	return false
}

func main() {
	// Define runtime environment
	// Default: development
	var config = appConfig{}
	appEnv, isSet := os.LookupEnv("APP_ENV")
	if !isSet {
		// If not set, fall back to development
		config.Runtime = "development"
	} else {
		// Make sure the runtime defined is actually valid
		if config.validRuntime(appEnv) {
			config.Runtime = appEnv
		} else {
			config.Runtime = "development" // default fallback
		}
	}

	appPort, isSet := os.LookupEnv("APP_PORT")
	if !isSet {
		config.Port = 9000
	} else {
		appPort, err := strconv.Atoi(appPort)
		if err != nil || appPort < 1000 {
			log.Println("App port", appPort, "could not be converted to integer or is a root-only port (< 1000)")
			config.Port = 9000
		} else {
			config.Port = appPort
		}
	}

	log.Println("App runtime is", config.Runtime)
	log.Println("App port is", config.Port)

	db, err := pop.Connect(config.Runtime)
	if err != nil {
		log.Panic(err)
	}

	entry := models.Entry{Diastolic: 80, Systolic: 120, Heartrate: 85}
	verrs, err := db.ValidateAndSave(&entry)
	if verrs.Count() > 0 {
		log.Println(fmt.Sprintf("Error while saving: %s\n", verrs))
	}
	if err != nil {
		log.Panic(err)
	}
}
