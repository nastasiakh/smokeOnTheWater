package main

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/db"
	"smokeOnTheWater/internal/handlers/auth"
	"smokeOnTheWater/internal/handlers/controllers"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/services"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	db.Init()

	userRepository := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", auth.Login)
		authGroup.POST("/sign-up", auth.SignUp)
	}
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.GetAllUsers)
	}
	router.Run(":8080")
}
