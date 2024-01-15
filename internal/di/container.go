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
	roleRepository := repositories.NewRoleRepository(db.DB)

	userService := services.NewUserService(userRepository, roleRepository)
	authService := services.NewAuthService(userRepository)
	roleService := services.NewRoleService(roleRepository)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, userService)
	roleController := controllers.NewRoleController(roleService)

	return &Container{
		UserController: userController,
		RoleController: roleController,
		AuthController: authController,
	}
}
