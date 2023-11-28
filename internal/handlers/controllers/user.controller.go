package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
	"strconv"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := new(models.User)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	if err := c.userService.Create(user); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}
	ctx.JSON(201, nil)
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
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.userService.GetById(uint(userId))
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
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	user, err := c.userService.Update(uint(userId), newUser)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update user"})
		return
	}
	ctx.JSON(201, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.userService.Delete(uint(userId)); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(204, nil)
}
