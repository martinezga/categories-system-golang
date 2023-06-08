package models

type CategoryModel struct {
	ID string `sql:"type:varchar(16);column:id;primaryKey;"`
	Name string
	// CreatedAt    time.Time
	// UpdatedAt    time.Time
}

func (CategoryModel) TableName() string {
    return "categories.categories"
}
