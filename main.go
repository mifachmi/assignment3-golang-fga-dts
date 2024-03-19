package main

import (
	"asssignment2/db"
	"asssignment2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.StartDB()

	r := gin.Default()

	api := r.Group("/api")
	routers.ItemsRouter(api)
	routers.OrderRouter(api)

	r.Run(":8080")
}
