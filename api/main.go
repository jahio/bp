package main

import (
	"github.com/gobuffalo/pop"
	"jahio/bp/router"
	"jahio/bp/config"
	"log"
	"strconv"
)

func main() {
	// Define runtime environment
	// Default: development
	var cfg = config.AppConfig{}
	cfg.SetRuntime("development") // default = development
	cfg.SetPort(9000) // default = 9000

	// Set pop to debug logging if not in prod
	if cfg.Runtime != "production" {
		pop.Debug = true
	}

	db, err := pop.Connect(cfg.Runtime)
	if err != nil {
		log.Panic(err)
	}

	// Instantiate the router and bind
	r := router.SetupRouter(db, &cfg)
	r.Run("0.0.0.0:" + strconv.Itoa(cfg.Port))
}
