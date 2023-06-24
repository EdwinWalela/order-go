package handlers

import (
	"net/http"

	"github.com/edwinwalela/go-order/orders/service"
	"github.com/gin-gonic/gin"
)

func GetProductById(c *gin.Context) {
	service.GetProductById()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func GetProducts(c *gin.Context) {
	service.GetProducts()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func UpdateProduct(c *gin.Context) {
	service.UpdateProduct()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func DeleteProduct(c *gin.Context) {
	service.DeleteProduct()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func CreateProduct(c *gin.Context) {
	service.CreateProduct()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
