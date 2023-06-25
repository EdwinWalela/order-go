package main

import (
	"context"
	"fmt"

	c "github.com/edwinwalela/go-order/orders/config"
	"github.com/edwinwalela/go-order/orders/db"
	"github.com/edwinwalela/go-order/orders/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := c.LoadConfig()
	ctx := context.Background()

	handler := handlers.Handler{
		Router: r,
	}

	handler.Register()

	conn, err := db.Up(ctx, config)

	if err != nil {
		panic(err)
	}

	_ = conn

	r.Run(fmt.Sprintf(":%s", config.Port))
}
