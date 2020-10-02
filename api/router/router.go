package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	"jahio/bp/config"
	"jahio/bp/controllers"
)

func SetupRouter(db *pop.Connection, cfg *config.AppConfig) *gin.Engine {
	if cfg.Runtime == "production" {
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
