package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
	"strconv"
)

type CategoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (c *CategoryController) GetAllCategories(ctx *gin.Context) {
	categories, err := c.categoryService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive categories"})
		return
	}
	ctx.JSON(200, categories)
}

func (c *CategoryController) GetCategoryById(ctx *gin.Context) {
	categoryId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "Page not found"})
		return
	}

	category, err := c.categoryService.GetById(uint(categoryId))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive category"})
	}

	ctx.JSON(200, category)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var category models.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid category data"})
		return
	}

	_, err := c.categoryService.Create(&category)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create category"})
		return
	}

	ctx.JSON(201, nil)
}

func (c *CategoryController) UpdateCategory(ctx *gin.Context) {
	var newCategory models.Category

	if err := ctx.ShouldBindJSON(&newCategory); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid category data"})
		return
	}

	categoryId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := c.categoryService.Update(uint(categoryId), &newCategory)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update category"})
		return
	}
	ctx.JSON(201, category)
}

func (c *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid category ID"})
		return
	}

	if err := c.categoryService.Delete(uint(categoryId)); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete category"})
		return
	}
	ctx.JSON(204, nil)

}
