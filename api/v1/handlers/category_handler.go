package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/shared"
	"github.com/martinezga/categories-system-golang/api/v1/dtos"
	"github.com/martinezga/categories-system-golang/api/v1/services/categories"
)

type resource struct {
	service *categories.CategoryService
}

// RegisterHandlers sets up the routing of the HTTP handlers
func RegisterHandlers(router *gin.Engine, service *categories.CategoryService) {
	res := resource{service}
	v1 := router.Group("/api/v1")
	{
		v1.POST("/categories", res.CreateCategoryHandler)
		v1.GET("/categories/:id", res.GetCategoryHandler)
	}
}

// GetCategoryHandler
// @Summary      Show category
// @Description  get category by id
// @Tags         v1
// @Produce      json
// @Param        id  path  string  true  "Category ID" default(cat_asdfgw58jy6j)
// @Success      200 {object}  shared.BaseResponse{data=dtos.CategoryResponse}  "Successful response"
// @Failure		 404 {object}  shared.ErrorResponse  "Not found response"
// @Router       /api/v1/categories/{id} [get]
func (r resource) GetCategoryHandler(c *gin.Context) {
	id := c.Param("id")
	category, err := r.service.Get(id)

	if err != nil {
		c.JSON(shared.NotFoundError(err, c.Request.URL.Path))
		return
	}

	c.JSON(shared.SuccessResponse(category))
}

// CreateCategoryHandler
// @Summary      Create category
// @Description  create category
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        category  body  dtos.CategoryCreateRequest  true  "Category"
// @Success      200 {object}  shared.BaseResponse{data=dtos.CategoryResponse}  "Successful response"
// @Failure		 400 {object}  shared.ErrorResponse  "Bad request response"
// @Router       /api/v1/categories [post]
func (r resource) CreateCategoryHandler(c *gin.Context) {
	var dataReceived dtos.CategoryCreateRequest

	if err := c.ShouldBindJSON(&dataReceived); err != nil {
		c.JSON(shared.BadRequestError(err, c.Request.URL.Path))
		return
	}

	category, err := r.service.Create(dataReceived)
	if err != nil {
		c.JSON(shared.BadRequestError(err, c.Request.URL.Path))
		return
	}

	c.JSON(shared.SuccessResponse(category))
}
