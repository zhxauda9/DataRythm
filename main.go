package main

import (
	h "datarythm/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/orders", h.CreateOrder)
	router.GET("/api/orders", h.GetOrders)
	router.GET("/api/orders/:id", h.GetOrder)
	router.PUT("/api/orders/:id", h.UpdateOrder)
	router.DELETE("/api/orders/:id", h.DeleteOrder)

	router.Run("localhost:8080")
}
