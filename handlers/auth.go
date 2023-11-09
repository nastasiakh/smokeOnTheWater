package handlers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/db"
	"smokeOnTheWater/models"
)

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid login data"})
		return
	}

	var foundUser models.User
	if err := db.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&foundUser).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}

func SignUp(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": "Invalid login data"})
		return
	}

	var existingUser models.User
	if err := db.DB.Where("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		c.JSON(400, gin.H{"error": "User with this email already exists"})
		return
	}

	if err := db.DB.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful"})
}
