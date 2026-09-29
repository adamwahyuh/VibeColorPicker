package authcontroller

import (
	"net/http"

	"github.com/adamwahyuh/coloroutfit/src/database"
	"github.com/adamwahyuh/coloroutfit/src/models"
	loginrequest "github.com/adamwahyuh/coloroutfit/src/request/LoginRequest"
	"github.com/adamwahyuh/coloroutfit/src/response"
	loginresponse "github.com/adamwahyuh/coloroutfit/src/response/LoginResponse"
	userresponse "github.com/adamwahyuh/coloroutfit/src/response/UserResponse"
	"github.com/adamwahyuh/coloroutfit/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var request loginrequest.LoginRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.Error[loginresponse.LoginResponse](
				http.StatusBadRequest,
				"Username and password must be filled.",
			),
		)
		return
	}

	var user models.User

	if err := database.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			response.Error[loginresponse.LoginResponse](
				http.StatusUnauthorized,
				"Username or password incorect.",
			),
		)

		return
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(request.Password),
	); err != nil {
		ctx.JSON(
			http.StatusUnauthorized,
			response.Error[loginresponse.LoginResponse](http.StatusUnauthorized, "Username or password incorect"),
		)
		return
	}

	token, err := utils.GenerateToken(user.Id)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.Error[loginresponse.LoginResponse](http.StatusInternalServerError, "Failed generating token"),
		)

		return
	}

	data := loginresponse.LoginResponse{
		Token: token,
		User: userresponse.UserResponse{
			Id:       &user.Id,
			Name:     &user.Name,
			Username: &user.Username,
			Email:    &user.Email,
		},
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)

	ctx.JSON(
		http.StatusOK,
		response.Success(http.StatusOK, "Successfully login", &data),
	)
}

func CheckLogin(ctx *gin.Context) {

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(
			http.StatusUnauthorized,
			response.Error[any](
				http.StatusUnauthorized,
				"Unauthenticated.",
			),
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.Success(
			http.StatusOK,
			"You Logged in.",
			&user,
		),
	)

}

func Logout(ctx *gin.Context) {
	ctx.SetCookie(
		"Authorization",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	ctx.JSON(http.StatusOK, response.Success[any](
		http.StatusOK,
		"Successfully Logged out",
		nil,
	))
}
