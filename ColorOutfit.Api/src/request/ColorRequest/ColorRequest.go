package colorrequest

type CreateCategoryRequest struct {
	Hex string `json:"hex" binding:"required,min=3,max=100"`
}
type UpdateCategoryRequest struct {
	Hex string `json:"hex" binding:"required,min=3,max=100"`
}
