package controllers

import (
	"github.com/gin-gonic/gin"
	"smokeOnTheWater/internal/handlers/services"
	"smokeOnTheWater/internal/models"
	"strconv"
	"time"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var requestData models.OrderWithProducts
	requestData.Order.DateCreated = time.Now()
	requestData.Order.DateModified = time.Now()

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		ctx.JSON(400, gin.H{"error": "Failed to decode request body"})
		return
	}

	createdOrder, err := c.orderService.Create(&requestData.Order, requestData.OrderProducts)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(201, createdOrder)
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.orderService.GetAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive orders"})
		return
	}
	ctx.JSON(200, orders)
}

func (c *OrderController) GetOrderById(ctx *gin.Context) {
	orderId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.orderService.GetById(uint(orderId))
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to receive order"})
		return
	}

	ctx.JSON(200, order)
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var newOrder models.OrderWithProducts
	newOrder.Order.DateModified = time.Now()

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid order data"})
		return
	}
	orderId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}
	order, err := c.orderService.Update(uint(orderId), newOrder)

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}
	ctx.JSON(201, order)
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid order ID"})
		return
	}
	if err := c.orderService.Delete(uint(orderId)); err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}

	ctx.JSON(204, nil)
}
