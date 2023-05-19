package categories

import (
	"github.com/martinezga/categories-system-golang/api/v1/models"
	"github.com/martinezga/categories-system-golang/api/v1/repositories"
)

	func GetCategoryService(repo repositories.Repository, id string) models.CategoryModel {

	return repo.GetCategoryRepository(id)
}
