package di

import (
	"smokeOnTheWater/internal/db"
	"smokeOnTheWater/internal/handlers/controllers"
	"smokeOnTheWater/internal/handlers/repositories"
	"smokeOnTheWater/internal/handlers/services"
)

type Container struct {
	UserController       *controllers.UserController
	RoleController       *controllers.RoleController
	AuthController       *controllers.AuthController
	PermissionController *controllers.PermissionController
	CategoryController   *controllers.CategoryController
	ProductController    *controllers.ProductController
	OrderController      *controllers.OrderController
}

func BuildContainer() *Container {
	userRepository := repositories.NewUserRepository(db.DB)
	roleRepository := repositories.NewRoleRepository(db.DB)
	permissionRepository := repositories.NewPermissionRepository(db.DB)
	categoryRepository := repositories.NewCategoryRepository(db.DB)
	productRepository := repositories.NewProductRepository(db.DB)
	orderRepository := repositories.NewOrderRepository(db.DB)
	orderProductRepository := repositories.NewOrderProductRepository(db.DB)

	userService := services.NewUserService(userRepository, roleRepository)
	authService := services.NewAuthService(userRepository)
	roleService := services.NewRoleService(roleRepository, permissionRepository)
	permissionService := services.NewPermissionService(permissionRepository)
	categoryService := services.NewCategoryService(categoryRepository)
	productService := services.NewProductService(productRepository)
	productQuantityCalculatorService := services.NewQuantityCalculatorService(productRepository)
	orderService := services.NewOrderService(orderRepository, orderProductRepository, productQuantityCalculatorService)

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService, userService)
	roleController := controllers.NewRoleController(roleService)
	permissionController := controllers.NewPermissionController(permissionService)
	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	return &Container{
		UserController:       userController,
		RoleController:       roleController,
		AuthController:       authController,
		PermissionController: permissionController,
		CategoryController:   categoryController,
		ProductController:    productController,
		OrderController:      orderController,
	}
}
