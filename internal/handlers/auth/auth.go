package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"smokeOnTheWater/internal/db"
	"smokeOnTheWater/internal/models"
)

var secretKey = []byte("your-secret-key")

func generateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
	})
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid login data"})
		return
	}

	var foundUser models.User
	if err := db.DB.Where("email = ?", user.Email).First(&foundUser).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := generateToken(foundUser)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{"token": token, "message": "Login successful"})
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	newUser.Password = string(hashedPassword)

	if err := db.DB.Create(&newUser).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful"})
}
