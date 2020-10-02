package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	"jahio/bp/models"
	"net/http"
	"log"
)

func NewEntryController(db *pop.Connection, c *gin.Context) {
	entry := models.Entry{}
	err := c.ShouldBind(&entry)
	if err != nil {
		// Couldn't parse the JSON into the entry object -- bad request
		status := StatusMessage{Status: "Error", Message: err.Error()}
		c.JSON(http.StatusBadRequest, status)
		log.Println(err)
		return
	}

	// Attempt to save the object to the database
	verrs, err := db.ValidateAndSave(&entry)
	if verrs.Count() > 0 {
		// Convert the string map to a singular string
		statusMsg := getValidationErrors(verrs.Errors)
		status := StatusMessage{Status: "Error", Message: statusMsg}
		c.JSON(http.StatusBadRequest, status)
		log.Println(statusMsg)
		return
	}
	if err != nil {
		status := StatusMessage{Status: "Error", Message: err.Error()}
		c.JSON(http.StatusInternalServerError, status)
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, entry)
}

func GetEntriesController(db *pop.Connection, c *gin.Context) {
	entries := []models.Entry{}
	err := db.Where("created_at > ?", c.Param("from")).Where("created_at < ?", c.Param("to")).All(&entries)
	if err != nil {
		status := StatusMessage{Status: "Error", Message: err.Error()}
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	c.JSON(http.StatusOK, entries)
}

func GetAllEntriesController(db *pop.Connection, c *gin.Context) {
	entries := []models.Entry{}
	err := db.All(&entries)
	if err != nil {
		status := StatusMessage{Status: "Error", Message: err.Error()}
		c.JSON(http.StatusInternalServerError, status)
		return
	}
	c.JSON(http.StatusOK, entries)
}
