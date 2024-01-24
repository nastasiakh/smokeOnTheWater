package routes

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/di"
)

func AddRoutes(router *gin.Engine, container *di.Container) {
	authController := container.AuthController
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", authController.Login)
		authGroup.POST("/sign-up", authController.SignUp)
	}

	userController := container.UserController
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userController.GetAllUsers)
		userGroup.GET("/:id", userController.GetUser)
		userGroup.POST("/", userController.CreateUser)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}

	roleController := container.RoleController
	roleGroup := router.Group("/roles")
	{
		roleGroup.GET("/", roleController.GetAllRoles)
		roleGroup.GET("/:id", roleController.GetRole)
		roleGroup.POST("/", roleController.CreateRole)
		roleGroup.PUT("/:id", roleController.UpdateRole)
		roleGroup.DELETE("/:id", roleController.DeleteRole)
	}

	permissionController := container.PermissionController
	permissionGroup := router.Group("/permissions")
	{
		permissionGroup.GET("/", permissionController.GetAllPermissions)
	}

	categoryController := container.CategoryController
	categoryGroup := router.Group("/categories")
	{
		categoryGroup.GET("/", categoryController.GetAllCategories)
		categoryGroup.GET("/:id", categoryController.GetCategoryById)
		categoryGroup.POST("/", categoryController.CreateCategory)
		categoryGroup.PUT("/:id", categoryController.UpdateCategory)
		categoryGroup.DELETE("/:id", categoryController.DeleteCategory)
	}

}
