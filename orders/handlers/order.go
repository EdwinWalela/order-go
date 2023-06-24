package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetOrders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func UpdateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func DeleteOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func CreateOrder(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
