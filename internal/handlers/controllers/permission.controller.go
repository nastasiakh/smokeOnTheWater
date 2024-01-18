package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
)

type PermissionController struct {
	permissionService *services.PermissionService
}

func NewPermissionController(permissionService *services.PermissionService) *PermissionController {
	return &PermissionController{permissionService: permissionService}
}

func (c *PermissionController) GetAllPermissions(ctx *gin.Context) {
	permissions, err := c.permissionService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive permissions"})
		return
	}
	ctx.JSON(200, permissions)
}
