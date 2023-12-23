package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"smokeOnTheWater/internal/db"
	"smokeOnTheWater/internal/db/migrations"
	"smokeOnTheWater/internal/di"
	"smokeOnTheWater/internal/handlers/middlewars"
	"smokeOnTheWater/internal/handlers/validation"
	"smokeOnTheWater/internal/routes"
)

func main() {
	validation.InitValidator()
	router := gin.Default()
	router.Use(middlewars.CorsMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	db.Init()
	if err := migrations.MigrateDB(db.DB); err != nil {
		panic("Failed to migrate database")
	}

	container := di.BuildContainer()

	routes.AddRoutes(router, container)

	router.Run(":8080")
}
