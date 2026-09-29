package usercontroller

import (
	"net/http"

	"github.com/adamwahyuh/coloroutfit/src/database"
	"github.com/adamwahyuh/coloroutfit/src/models"
	userrequest "github.com/adamwahyuh/coloroutfit/src/request/UserRequest"
	"github.com/adamwahyuh/coloroutfit/src/response"
	userresponse "github.com/adamwahyuh/coloroutfit/src/response/UserResponse"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func StoreUser(ctx *gin.Context) {
	var request userrequest.CreateUserRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.Error[userresponse.UserResponse](
				http.StatusBadRequest,
				"Invalid Request",
			),
		)

		return
	}

	if request.Password != request.PasswordConfirm {
		ctx.JSON(
			http.StatusBadRequest,
			response.Error[userresponse.UserResponse](
				http.StatusBadRequest,
				"Password do not mathc",
			),
		)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.Error[userresponse.UserResponse](
				http.StatusInternalServerError,
				"Failed to proccess password",
			),
		)
		return
	}

	password := string(hashedPassword)

	user := models.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: password,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.Error[userresponse.UserResponse](
				http.StatusInternalServerError,
				"Failed creating user.",
			),
		)
		return
	}

	data := userresponse.UserResponse{
		Id:       &user.Id,
		Name:     &user.Name,
		Username: &user.Username,
		Email:    &user.Email,
	}

	ctx.JSON(
		http.StatusCreated,
		response.Success(http.StatusCreated, "User successfully created", &data),
	)
}
