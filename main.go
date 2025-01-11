package main

import (
	"net/http"
	"stopit/controllers"
	"stopit/middleware"
	"stopit/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()
	auth := r.Group("/")

	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/action", controllers.AllAction)
	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Written in go by tb12as"})
	})
	r.POST("/login", controllers.Login)

	r.Run(":8000")
}
