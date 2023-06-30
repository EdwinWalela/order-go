package customer

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Customer struct {
	Handler Handler
	Router  *gin.Engine
	DBConn  *firestore.Client
	Context context.Context
	Logger  *zap.Logger
}

func (c *Customer) Register() {
	c.Handler.service.repository.DBConn = c.DBConn
	c.Handler.service.repository.context = c.Context
	c.Handler.logger = c.Logger
	c.Handler.service.repository.logger = c.Logger

	v1 := c.Router.Group("/api/v1")
	v1.GET("/customers/:id", c.Handler.GetById)
	v1.GET("/customers", c.Handler.GetAll)
	v1.DELETE("/customers/:id", c.Handler.Delete)
	v1.POST("/customers", c.Handler.Create)
}
