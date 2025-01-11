package controllers

import (
	"net/http"
	"stopit/middleware"
	"stopit/models"

	"github.com/gin-gonic/gin"
)

func AllAction(c *gin.Context) {
	user := middleware.User

	var actions []models.Action
	models.DB.Where("user_id = ?", user.Id).Find(&actions)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		// "user":    user,
		"data": actions,
	})
}

func CreateAction(c *gin.Context) {
	// TODO : implement
}
