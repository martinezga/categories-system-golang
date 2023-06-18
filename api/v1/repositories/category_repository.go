package repositories

import (
	"github.com/martinezga/categories-system-golang/api/v1/dtos"
	"github.com/martinezga/categories-system-golang/api/v1/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetCategoryRepository(id string) (models.CategoryModel, error)
	CreateCategoryRepository(data dtos.CategoryCreateRequest) (models.CategoryModel, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db}
}

func (r repository) GetCategoryRepository(id string) (models.CategoryModel, error) {
	var category models.CategoryModel
	err := r.db.First(&category, models.CategoryModel{ID: id}).Error

	return category, err
}

func (r repository) CreateCategoryRepository(data dtos.CategoryCreateRequest) (models.CategoryModel, error) {
	var category models.CategoryModel
	category.Name = data.Name
	category.Description = data.Description

	err := r.db.Create(&category).Error

	return category, err
}
