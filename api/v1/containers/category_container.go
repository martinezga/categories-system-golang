package containers

import (
	"github.com/gin-gonic/gin"
	"github.com/martinezga/categories-system-golang/api/v1/handlers"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
	"github.com/martinezga/categories-system-golang/api/v1/services/categories"
	"gorm.io/gorm"
)

type CategoryContainer struct {
	CategoryHandler gin.HandlerFunc
}

func BuildCategoryContainer(router *gin.Engine, db *gorm.DB) *CategoryContainer {
	categoryRepository := repositories.NewRepository(db)
	categoryService := categories.NewService(categoryRepository)
	handlers.RegisterHandlers(router, categoryService)

	return &CategoryContainer{
		CategoryHandler: router.HandleContext,
	}
}
