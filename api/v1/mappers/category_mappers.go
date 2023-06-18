package mappers

import (
	"github.com/martinezga/categories-system-golang/api/v1/dtos"
	"github.com/martinezga/categories-system-golang/api/v1/models"
)

func MapToCategoryResponse(category models.CategoryModel) dtos.CategoryResponse {
	return dtos.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}
