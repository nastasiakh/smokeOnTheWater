package di

import (
	"smokeOnTheWater/internal/db"
	"smokeOnTheWater/internal/handlers/controllers"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/services"
)

type Container struct {
	UserController *controllers.UserController
	RoleController *controllers.RoleController
	AuthController *controllers.AuthController
}

func BuildContainer() *Container {
	userRepository := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	roleRepository := repositories.NewRoleRepository(db.DB)
	roleService := services.NewRoleService(roleRepository)
	roleController := controllers.NewRoleController(roleService)

	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	return &Container{
		UserController: userController,
		RoleController: roleController,
		AuthController: authController,
	}
}
