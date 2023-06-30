package customer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Handler struct {
	service Service
	logger  *zap.Logger
}

func (h *Handler) GetById(c *gin.Context) {
	id := c.Param("id")
	h.logger.Info("Fetching customer by uuid", zap.String("uuid", id))
	customer, err := h.service.GetById(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
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
	h.logger.Info("Fetching all customers")
	customers, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"customers": customers,
	})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	h.logger.Info("Deleting customer", zap.String("uuid", id))
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
		"id":      id,
	})
}

func (h *Handler) Create(c *gin.Context) {
	h.logger.Info("Creating customer")
	jsonBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	model := Model{}
	err = json.Unmarshal(jsonBody, &model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uuid, err := h.service.Create(model)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.logger.Info("Customer created", zap.String("uuid", uuid))
	c.JSON(http.StatusCreated, gin.H{
		"message": "customer created",
		"uuid":    uuid,
	})
}
