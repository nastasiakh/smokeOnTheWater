package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
	"strconv"
)

type ProductController struct {
	productService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product data"})
		return
	}

	_, err := c.productService.Create(&product)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create product"})
		return
	}
	ctx.JSON(201, nil)
}

func (c *ProductController) GetAllProducts(ctx *gin.Context) {
	products, err := c.productService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive products"})
		return
	}
	ctx.JSON(200, products)
}

func (c *ProductController) GetProduct(ctx *gin.Context) {
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := c.productService.GetById(uint(productId))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive product"})
		return
	}
	ctx.JSON(200, product)
}

func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var newProduct models.Product
	if err := ctx.ShouldBindJSON(&newProduct); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product data"})
		return
	}
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := c.productService.Update(uint(productId), &newProduct)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update product"})
		return
	}
	ctx.JSON(201, product)
}

func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	productId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := c.productService.Delete(uint(productId)); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete product"})
		return
	}
	ctx.JSON(204, nil)
}
