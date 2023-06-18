package models

import (
	"github.com/martinezga/categories-system-golang/api/shared"
	"gorm.io/gorm"
)

type CategoryModel struct {
	ID          string `sql:"type:varchar(16);column:id;primaryKey;"`
	Name        string
	Description string
}

func (CategoryModel) TableName() string {
	return "categories.categories"
}

func (category *CategoryModel) BeforeCreate(*gorm.DB) (err error) {
	id, _ := shared.GenerateUniqueString("cat", 16)
	category.ID = id
	return
}
