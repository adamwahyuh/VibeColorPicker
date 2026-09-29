package userresponse

type UserResponse struct {
	Id       *uint16 `json:"id"`
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}
