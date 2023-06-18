package categories

import (
	"github.com/martinezga/categories-system-golang/api/v1/dtos"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
)

// Service encapsulates usecase logic
type CategoryService struct {
	Create func(category dtos.CategoryCreateRequest) (dtos.CategoryResponse, error)
	Get    func(id string) (dtos.CategoryResponse, error)
}

type service struct {
	repo repositories.Repository
}

func NewService(repo repositories.Repository) *CategoryService {
	return &CategoryService{
		Create: NewCreateService(repo).CreateCategory,
		Get:    NewGetService(repo).GetCategory,
	}
}
