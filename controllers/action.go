package controllers

import (
	"net/http"
	"stopit/middleware"
	"stopit/models"
	"time"

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
	var request models.CreateAction
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(422, gin.H{"message": "Failed to handle request", "success": false})
		return
	}

	action := models.Action{
		Name:      request.Name,
		UserId:    middleware.User.Id,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	result := models.DB.Create(&action)
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusCreated, map[string]any{"message": "Action created", "action": action})
}
