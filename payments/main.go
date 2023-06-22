package main

import (
	"fmt"

	c "github.com/edwinwalela/go-order/payments/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := c.LoadConfig()

	r.Run(fmt.Sprintf(":%s", config.Port))
}
