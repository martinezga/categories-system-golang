package categories

import (
	"github.com/martinezga/categories-system-golang/api/v1/models"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
)

// Service encapsulates usecase logic
type CategoryService struct {
	Get func(id string) (models.CategoryModel, error)
}

type service struct {
	repo   repositories.Repository
}

func NewService(repo repositories.Repository) *CategoryService {
	return &CategoryService{
		Get:    NewGetService(repo).GetCategory,
		// Create: NewCreateService(repo),
	}
}
