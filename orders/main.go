package main

import (
	"context"
	"fmt"

	c "github.com/edwinwalela/go-order/orders/config"
	"github.com/edwinwalela/go-order/orders/db"
	"github.com/edwinwalela/go-order/orders/service/customer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := gin.Default()
	config := c.LoadConfig()
	ctx := context.Background()

	conn, err := db.Up(ctx, config)

	if err != nil {
		panic(err)
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	customers := customer.Customer{
		Router:  r,
		DBConn:  conn,
		Context: ctx,
		Logger:  logger,
	}
	customers.Register()

	r.Run(fmt.Sprintf(":%s", config.Port))
}
