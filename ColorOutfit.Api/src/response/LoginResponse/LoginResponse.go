package loginresponse

import userresponse "github.com/adamwahyuh/coloroutfit/src/response/UserResponse"

type LoginResponse struct {
	Token string                    `json:"token"`
	User  userresponse.UserResponse `json:"user"`
}
