package templaterequest

type CreateInitializeTemplate struct {
	Title       string `json:"name" binding:"required,min=3,max=50"`
	Description string `json:"username" binding:"max=255"`
}

type AddColorToTemplate struct {
	ColorId uint64 `json:"color_id" binding:"required,gt=0"`
}
