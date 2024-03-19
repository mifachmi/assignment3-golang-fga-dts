package routers

import (
	"asssignment2/controllers"

	"github.com/gin-gonic/gin"
)

func ItemsRouter(r *gin.RouterGroup) {
	r.GET("/items", controllers.GetItems)
	r.POST("/items", controllers.CreateItem)
}
