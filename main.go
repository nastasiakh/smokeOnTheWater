package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"smokeOnTheWater/internal/db"
	"smokeOnTheWater/internal/handlers/controllers"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/handlers/validation"
)

func main() {
	validation.InitValidator()
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	db.Init()
	if err := db.MigrateDB(db.DB); err != nil {
		panic("Failed to migrate database")
	}

	userRepository := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(services.NewAuthService(userRepository))

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/sign-up", authController.SignUp)
	}
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
	router.Run(":8080")
}
