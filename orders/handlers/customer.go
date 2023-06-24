package handlers

import (
	"net/http"

	"github.com/edwinwalela/go-order/orders/service"
	"github.com/gin-gonic/gin"
)

func GetCustomerById(c *gin.Context) {
	service.CreateCustomer()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetCustomers(c *gin.Context) {
	service.GetCustomers()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func UpdateCustomer(c *gin.Context) {
	service.UpdateCustomer()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func DeleteCustomer(c *gin.Context) {
	service.DeleteCustomer()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func CreateCustomer(c *gin.Context) {
	service.CreateCustomer()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
