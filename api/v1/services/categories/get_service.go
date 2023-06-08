package categories

import (
	"github.com/martinezga/categories-system-golang/api/v1/models"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
)

type GetService struct {
	repo repositories.Repository
}

func NewGetService(repo repositories.Repository) *GetService {
	return &GetService{repo: repo}
}

func (s *GetService) GetCategory(id string) (models.CategoryModel, error) {
	category, err := s.repo.GetCategoryRepository(id)
	// return mapToCategoryResponse(category)
	if err != nil {
		return category, err
	}

	return category, nil
}
