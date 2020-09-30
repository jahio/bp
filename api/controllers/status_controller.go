package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	"net/http"
)

func StatusController(db *pop.Connection, c *gin.Context) {
	status := StatusMessage{Status: "OK", Message: "All systems go!"}
	c.JSON(http.StatusOK, status)
}
