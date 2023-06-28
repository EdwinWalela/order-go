package customer

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.service.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customer": customer,
	})
}

func (h *Handler) GetAll(c *gin.Context) {
	h.service.GetAll()
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *Handler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (h *Handler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
