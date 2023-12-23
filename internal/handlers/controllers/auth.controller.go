package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
)

var secretKey = []byte("your-secret-key")

func generateToken(user *models.User) (string, error) {
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

type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewAuthController(authService *services.AuthService, userService *services.UserService) *AuthController {
	return &AuthController{authService: authService, userService: userService}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid login data"})
		return
	}

	existingUser, err := c.authService.Authenticate(loginData.Email, loginData.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := generateToken(existingUser)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "message": "Login successful"})
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var newUser models.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid register data"})
		return
	}

	existingUser, err := c.userService.GetByEmail(newUser.Email)
	if existingUser != nil {
		ctx.JSON(400, gin.H{"error": "User with such email already exists"})
		return
	}

	createdUser, err := c.userService.Create(&newUser)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	token, err := generateToken(createdUser)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "message": "Registration successful"})
}
