package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	if err := c.userService.Create(&user); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive users"})
		return
	}
	ctx.JSON(200, users)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	var userId uint
	if err := ctx.ShouldBindJSON(&userId); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}
	user, err := c.userService.GetById(userId)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive user"})
		return
	}
	ctx.JSON(200, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}
	user, err := c.userService.Update(newUser.ID, newUser)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}
	ctx.JSON(200, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	var userId uint
	if err := ctx.ShouldBindJSON(&userId); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	if err := c.userService.Delete(userId); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(200, nil)
}
