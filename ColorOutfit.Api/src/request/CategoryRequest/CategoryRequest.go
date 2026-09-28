package categoryrequest

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
type UpdateCategoryRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
