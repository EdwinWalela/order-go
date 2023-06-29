package customer

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	Handler Handler
	Router  *gin.Engine
	DBConn  *firestore.Client
	Context context.Context
}

func (c *Customer) Register() {
	c.Handler.service.repository.DBConn = c.DBConn
	c.Handler.service.repository.context = c.Context

	v1 := c.Router.Group("/api/v1")
	v1.GET("/customers/:id", c.Handler.GetById)
	v1.GET("/customers", c.Handler.GetAll)
	v1.PATCH("/customers/:id", c.Handler.Update)
	v1.DELETE("/customers/:id", c.Handler.Delete)
	v1.POST("/customers", c.Handler.Create)
}
