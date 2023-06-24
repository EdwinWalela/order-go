package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
	Router *gin.Engine
}

func (h *Handler) Register() {
	h.Router.GET("/orders/:id", GetOrderById)
	h.Router.GET("/orders", GetOrders)
	h.Router.PATCH("/orders/:id", UpdateOrder)
	h.Router.DELETE("/orders/:id", DeleteOrder)
	h.Router.POST("/orders", CreateOrder)

	h.Router.GET("/customers/:id", GetCustomerById)
	h.Router.GET("/customers", GetCustomers)
	h.Router.PATCH("/customers/:id", UpdateCustomer)
	h.Router.DELETE("/customers/:id", DeleteCustomer)
	h.Router.POST("/customers", CreateCustomer)

}
