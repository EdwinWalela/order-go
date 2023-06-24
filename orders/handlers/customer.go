package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomerById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func UpdateCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func DeleteCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func CreateCustomer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
