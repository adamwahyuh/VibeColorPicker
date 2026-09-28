package userrequest

type CreateUserRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=100"`
	Username        string `json:"username" binding:"required,alphanum,lowercase,min=4,max=20"`
	Email           string `json:"email" binding:"required,email,max=255"`
	Password        string `json:"password" binding:"required,min=8,max=64"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=8,max=64"`
}

type ChangeUsernameRequest struct {
	Username string `json:"username" binding:"required,alphanum,lowercase,min=4,max=20"`
}

type ChangeNameRequest struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
