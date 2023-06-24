package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
	Router *gin.Engine
}

func (h *Handler) Register() {
	v1 := h.Router.Group("/api/v1")

	v1.GET("/orders/:id", GetOrderById)
	v1.GET("/orders", GetOrders)
	v1.PATCH("/orders/:id", UpdateOrder)
	v1.DELETE("/orders/:id", DeleteOrder)
	v1.POST("/orders", CreateOrder)

	v1.GET("/customers/:id", GetCustomerById)
	v1.GET("/customers", GetCustomers)
	v1.PATCH("/customers/:id", UpdateCustomer)
	v1.DELETE("/customers/:id", DeleteCustomer)
	v1.POST("/customers", CreateCustomer)

	v1.GET("/products/:id", GetProductById)
	v1.GET("/products", GetProducts)
	v1.PATCH("/products/:id", UpdateProduct)
	v1.DELETE("/products/:id", DeleteProduct)
	v1.POST("/products", CreateProduct)

}
