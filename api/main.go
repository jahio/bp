package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	"jahio/bp/config"
	"jahio/bp/router"
	"log"
	"os"
	"strconv"
)

func main() {
	var cfg = config.AppConfig{}

	// Figure out the current runtime environment; default is development
	var rt = "development" // default
	switch os.Getenv("APP_ENV") {
	case "test":
		rt = "test"
	case "production":
		rt = "production"
	default:
		rt = "development"
	}

	// Figure out the API port from environment or set it (default: 9000)
	apiPort, _ := strconv.Atoi(os.Getenv("APP_PORT"))
	if apiPort < 1000 {
		fmt.Println("Erroneous or no API port specified, starting API on port 9000")
		apiPort = 9000
	}

	cfg.SetRuntime(rt) // default = development
	cfg.SetPort(apiPort) // default = 9000

	// Set pop to debug logging if not in prod
	pop.Debug = true

	if cfg.Runtime == "production" {
		pop.Debug = false
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := pop.Connect(cfg.Runtime)
	if err != nil {
		log.Panic(err)
	}

	// Instantiate the router and bind
	r := router.SetupRouter(db, &cfg)
	r.Run("0.0.0.0:" + strconv.Itoa(cfg.Port))
}
