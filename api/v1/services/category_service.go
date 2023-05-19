package services

import (
	"github.com/martinezga/categories-system-golang/api/v1/models"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
	"github.com/martinezga/categories-system-golang/api/v1/services/categories"
)

// Service encapsulates usecase logic
type Service interface {
	GetCategoryService(id string) models.CategoryModel
}

type service struct {
	repo   repositories.Repository
}

func NewService(repo repositories.Repository) Service {
	return service{repo}
}

func (s service) GetCategoryService(id string) models.CategoryModel {
	return categories.GetCategoryService(s.repo, id)
}
