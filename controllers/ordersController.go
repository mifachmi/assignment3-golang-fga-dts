package controllers

import (
	"asssignment2/db"
	"asssignment2/helpers"
	"asssignment2/models"
	"asssignment2/requests"
	"asssignment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var orderService services.OrderService

func GetOrder(ctx *gin.Context) {
	var orders []models.Order
	db.GetDB().Preload("Items").Find(&orders)

	ctx.JSON(http.StatusOK, orders)
}

func GetOrderById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{Message: "invalid required param"})
		return
	}

	var orderWithItems models.Order
	db.GetDB().Preload("Items").First(&orderWithItems, "order_id = ?", id)

	result := map[string]interface{}{
		"customerName": orderWithItems.CustomerName,
		"orderedAt":    orderWithItems.OrderedAt,
		"items":        orderWithItems.Items,
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateOrder(ctx *gin.Context) {
	var orderWithItems requests.OrderWithItemsBody

	if err := ctx.BindJSON(&orderWithItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := orderService.CreateOrder(orderWithItems)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Order data not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Success create data order", "data": order})
}

func UpdateOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))

	var orderWithItems requests.OrderWithItemsBody

	if orderId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{Message: "invalid required param"})
		return
	}

	if err := ctx.ShouldBindJSON(&orderWithItems); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := orderService.UpdateOrder(uint(orderId), orderWithItems)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Order data not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Success update data order", "data": order})
}

func DeleteOrder(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderId"))

	if orderId == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{Message: "Invalid required param"})
		return
	}

	if err := orderService.DeleteOrder(uint(orderId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order and associated items deleted successfully"})
}
