package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
	"strconv"
)

type RoleController struct {
	roleService *services.RoleService
}

func NewRoleController(roleService *services.RoleService) *RoleController {
	return &RoleController{roleService: roleService}
}

func (c *RoleController) GetAllRoles(ctx *gin.Context) {
	roles, err := c.roleService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive roles"})
		return
	}

	ctx.JSON(200, roles)
}

func (c *RoleController) GetRole(ctx *gin.Context) {
	roleId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := c.roleService.GetById(uint(roleId))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive role"})
		return
	}
	ctx.JSON(200, role)
}

func (c *RoleController) CreateRole(ctx *gin.Context) {
	var role models.Role
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid role data"})
		return
	}

	if err := c.roleService.Create(role); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create role"})
		return
	}
	ctx.JSON(201, nil)
}

func (c *RoleController) UpdateRole(ctx *gin.Context) {
	var newRole models.Role
	if err := ctx.ShouldBindJSON(&newRole); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to update user"})
		return
	}
	roleId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid role ID"})
		return
	}
	role, err := c.roleService.Update(uint(roleId), newRole)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update role"})
		return
	}
	ctx.JSON(300, role)
}

func (c *RoleController) DeleteRole(ctx *gin.Context) {
	roleId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := c.roleService.Delete(uint(roleId)); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete role"})
		return
	}
	ctx.JSON(204, nil)
}
