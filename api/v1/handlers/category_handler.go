package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/services/categories"
)

type resource struct {
	service *categories.CategoryService
}

// RegisterHandlers sets up the routing of the HTTP handlers
func RegisterHandlers(router *gin.Engine, service *categories.CategoryService) {
	res := resource{service}

	router.GET("/v1/categories/:id", res.GetCategoryHandler)
}

func (r resource) GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := r.service.Get(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": category,
	})
}
