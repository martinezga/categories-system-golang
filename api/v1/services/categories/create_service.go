package categories

import (
	"github.com/martinezga/categories-system-golang/api/v1/dtos"
	"github.com/martinezga/categories-system-golang/api/v1/mappers"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
)

type CreateService struct {
	repo repositories.Repository
}

func NewCreateService(repo repositories.Repository) *CreateService {
	return &CreateService{repo: repo}
}

func (s *CreateService) CreateCategory(data dtos.CategoryCreateRequest) (dtos.CategoryResponse, error) {
	category, err := s.repo.CreateCategoryRepository(data)

	if err != nil {
		return dtos.CategoryResponse{}, err
	}

	return mappers.MapToCategoryResponse(category), nil
}
