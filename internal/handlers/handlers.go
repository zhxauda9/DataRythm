package handlers

import (
	o "datarythm/internal/order"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	orders []o.Order
	nextID = 1
	mu     sync.Mutex
)

func CreateOrder(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var newOrder o.Order
	if err := c.ShouldBindJSON(&newOrder); err == nil {
		newOrder.ID = nextID
		nextID++
		orders = append(orders, newOrder)
		c.JSON(http.StatusCreated, newOrder)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func GetOrders(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	for _, order := range orders {
		if fmt.Sprint(order.ID) == id {
			c.JSON(http.StatusOK, order)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}

func UpdateOrder(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	var updatedOrder o.Order
	if err := c.ShouldBindJSON(&updatedOrder); err == nil {
		for i, order := range orders {
			if fmt.Sprint(order.ID) == id {
				orders[i].Product = updatedOrder.Product
				orders[i].Quantity = updatedOrder.Quantity
				c.JSON(http.StatusOK, orders[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func DeleteOrder(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id := c.Param("id")
	for i, order := range orders {
		if fmt.Sprint(order.ID) == id {
			orders = append(orders[:i], orders[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
}
