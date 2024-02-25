package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

const (
	accessExpire  = 15 * time.Minute
	refreshExpire = 24 * time.Hour
)

func generateTokens(user *models.User) (accessToken, refreshToken string, err error) {
	accessClaims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(accessExpire).Unix(),
	}

	refreshClaims := jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(refreshExpire).Unix(),
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %v", err)
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %v", err)
	}
	return accessToken, refreshToken, nil
}

type AuthController struct {
	authService *services.AuthService
	userService *services.UserService
}

func NewAuthController(authService *services.AuthService, userService *services.UserService) *AuthController {
	return &AuthController{authService: authService, userService: userService}
}

func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var refreshData struct {
		RefreshToken string `json:"refreshToken" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&refreshData); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid refresh token data"})
		return
	}

	token, err := jwt.Parse(refreshData.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid refresh token"})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if _, ok := claims["exp"].(float64); !ok {
			ctx.JSON(401, gin.H{"error": "Invalid token expiration"})
			return
		}

		userID := claims["userID"].(uint)
		user, err := c.userService.GetById(userID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to get user"})
			return
		}

		accessToken, _, err := generateTokens(&user)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to generate access token"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
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
	ctx.Set("userRole", existingUser.Roles)
	token, _, err := generateTokens(existingUser)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "user": existingUser, "message": "Login successful"})
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

	token, _, err := generateTokens(createdUser)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(200, gin.H{"token": token, "message": "Registration successful"})
}
