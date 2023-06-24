package handlers

import (
	"net/http"

	"github.com/edwinwalela/go-order/orders/service"
	"github.com/gin-gonic/gin"
)

func GetOrderById(c *gin.Context) {
	service.GetOrderById()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetOrders(c *gin.Context) {
	service.GetOrders()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func UpdateOrder(c *gin.Context) {
	service.UpdateOrder()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func DeleteOrder(c *gin.Context) {
	service.DeleteOrder()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func CreateOrder(c *gin.Context) {
	service.CreateOrder()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
