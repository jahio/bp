package controllers

import(
	_ "jahio/bp/models"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	"net/http"
)

type statusMessage struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func StatusController(db *pop.Connection, c *gin.Context) {
	status := statusMessage{Status:"OK", Message:"All systems go!"}
	c.JSON(http.StatusOK, status)
}
