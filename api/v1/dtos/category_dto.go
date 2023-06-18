package dtos

type CategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryCreateRequest struct {
	Name        string `json:"name" binding:"required" default:"test"`
	Description string `json:"description" default:"test description"`
}
