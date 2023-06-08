package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/services"
)

type resource struct {
	service services.Service
}

// RegisterHandlers sets up the routing of the HTTP handlers
func RegisterHandlers(router *gin.Engine, service services.Service) {
	res := resource{service}

	router.GET("/v1/category/:id", res.GetCategoryHandler)
}

func (r resource) GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	category := r.service.GetCategoryService(id)
	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
