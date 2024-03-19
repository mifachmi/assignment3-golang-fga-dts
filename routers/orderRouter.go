package routers

import (
	"asssignment2/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRouter(r *gin.RouterGroup) {
	r.GET("/orders", controllers.GetOrder)
	r.POST("/orders", controllers.CreateOrder)
	r.PUT("/orders/:orderId", controllers.UpdateOrder)
	r.DELETE("/orders/:orderId", controllers.DeleteOrder)
}
