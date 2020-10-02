package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	Port    int
	Runtime string
}

func (a *AppConfig) ValidRuntime(r string) bool {
	switch r {
		case
			"development", "test", "production":
			return true
	}
	return false
}

func (a *AppConfig) SetRuntime(def string) {
	appEnv, isSet := os.LookupEnv("APP_ENV")
	if !isSet {
		// If not set, fall back to default passed in
		a.Runtime = def
		log.Println("APP_ENV isn't set, using default:", a.Runtime)
		return
	}

	if !a.ValidRuntime(appEnv) {
		a.Runtime = def
		log.Println("APP_ENV isn't a valid runtime, using default:", a.Runtime)
		return
	}

	a.Runtime = appEnv
	log.Println("Using Runtime ENV", a.Runtime)
	return
}

func (a *AppConfig) SetPort(def int) {
	appPort, isSet := os.LookupEnv("APP_PORT")
	if !isSet {
		a.Port = def
		log.Println("No APP_PORT env set. Using default:", a.Port)
		return
	}

	// Convert it to an integer
	port, err := strconv.Atoi(appPort)
	if err != nil {
		a.Port = def
		log.Println("Error converting", port, "to integer, using default:", a.Port)
		return
	}

	// Check if using a root-only port
	if port < 1000 {
		a.Port = def
		log.Println("Error:", port, "is a root-only port (< 1000). Using default:", a.Port)
		return
	}

	// If we get this far, appPort is set, it's an int, and it's > 1000
	a.Port = port
	log.Println("Binding to port", a.Port)
	return
}
