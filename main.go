package main

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/db"
	"smokeOnTheWater/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	db.Init()
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", handlers.Login)
		authGroup.POST("/sign-up", handlers.SignUp)
	}

	router.Run(":8080")
}
