package controllers

import (
	"net/http"
	"stopit/models"

	"github.com/gin-gonic/gin"
)

func AllAction(c *gin.Context) {
	var actions []models.Action
	models.DB.Find(&actions)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    actions,
	})
}

func CreateAction(c *gin.Context) {
	// TODO : implement
}
