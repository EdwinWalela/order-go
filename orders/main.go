package main

import (
	"fmt"

	c "github.com/edwinwalela/go-order/orders/config"
	"github.com/edwinwalela/go-order/orders/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := c.LoadConfig()

	handler := handlers.Handler{
		Router: r,
	}

	handler.Register()

	r.Run(fmt.Sprintf(":%s", config.Port))
}
