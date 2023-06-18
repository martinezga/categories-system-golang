package categories

import (
	"github.com/martinezga/categories-system-golang/api/v1/dtos"
	"github.com/martinezga/categories-system-golang/api/v1/mappers"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
)

type GetService struct {
	repo repositories.Repository
}

func NewGetService(repo repositories.Repository) *GetService {
	return &GetService{repo: repo}
}

func (s *GetService) GetCategory(id string) (dtos.CategoryResponse, error) {
	category, err := s.repo.GetCategoryRepository(id)

	if err != nil {
		return dtos.CategoryResponse{}, err
	}

	return mappers.MapToCategoryResponse(category), nil
}
