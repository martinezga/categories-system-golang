package containers

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
	"github.com/martinezga/categories-system-golang/api/v1/services"
	"github.com/martinezga/categories-system-golang/api/v1/handlers"
)

type CategoryContainer struct {
	CategoryHandler gin.HandlerFunc
}

func BuildCategoryContainer(router *gin.Engine) *CategoryContainer {
	categoryRepository := repositories.NewRepository()
	categoryService := services.NewService(categoryRepository)
	handlers.RegisterHandlers(router, categoryService)

	return &CategoryContainer{
		CategoryHandler: router.HandleContext,
	}
}
