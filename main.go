package main

import (
	"stopit/controllers"
	"stopit/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	r.GET("/action", controllers.AllAction)

	r.Run(":8000")
}
