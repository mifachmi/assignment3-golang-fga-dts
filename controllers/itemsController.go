package controllers

import (
	"asssignment2/db"
	"asssignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(ctx *gin.Context) {
	var items []models.Item
	db.GetDB().Find(&items)
	ctx.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.GetDB().Create(&newItem)
	c.JSON(http.StatusCreated, newItem)
}
