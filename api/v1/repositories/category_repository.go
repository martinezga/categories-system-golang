package repositories

import (
	"github.com/martinezga/categories-system-golang/api/v1/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetCategoryRepository(id string) (models.CategoryModel, error)
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
