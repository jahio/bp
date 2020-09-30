package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	"jahio/bp/controllers"
	"log"
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

func setupRouter(db *pop.Connection, config *appConfig) *gin.Engine {
	if config.Runtime == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		controllers.StatusController(db, c)
	})

	r.GET("/entries", func(c *gin.Context) {
		controllers.GetAllEntriesController(db, c)
	})

	r.GET("/entries/:from/:to", func(c *gin.Context) {
		controllers.GetEntriesController(db, c)
	})

	r.POST("/entries", func(c *gin.Context) {
		controllers.NewEntryController(db, c)
	})

	return r
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

	// Set the application port
	// Default fallback: 9000
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

	log.Println("Starting app in", config.Runtime, "mode")
	log.Println("Binding to port", config.Port)

	// Set pop to debug logging if not in prod
	if config.Runtime != "production" {
		pop.Debug = true
	}

	db, err := pop.Connect(config.Runtime)
	if err != nil {
		log.Panic(err)
	}

	// Instantiate the router and bind
	r := setupRouter(db, &config)
	r.Run("0.0.0.0:" + strconv.Itoa(config.Port))
}
