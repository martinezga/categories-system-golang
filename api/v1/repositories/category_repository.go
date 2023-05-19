package repositories

import (
	"fmt"

	"github.com/martinezga/categories-system-golang/api/v1/models"
)

type Repository interface {
	GetCategoryRepository(id string) models.CategoryModel
}

type repository struct {
	// TODO: db
}

func NewRepository() Repository {
	return repository{}
}

func (r repository) GetCategoryRepository(id string) models.CategoryModel {
	var category models.CategoryModel
	category = models.CategoryModel{}
	category.Name = fmt.Sprintf("Category %s", id)

	return category
}
