package main

import (
	"net/http"
	"stopit/config"
	"stopit/controllers"
	"stopit/middleware"
	"stopit/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitEnv()
	models.ConnectDatabase()

	r := gin.Default()
	auth := r.Group("/")

	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/action", controllers.AllAction)
		auth.POST("/action", controllers.CreateAction)
	}

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Written in go by tb12as"})
	})
	r.POST("/login", controllers.Login)

	port := config.Env("RUNNING_PORT")
	if port == "" {
		panic("No port was set in .env file")
	}
	r.Run(":" + port)
}
