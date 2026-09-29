package middleware

import (
	"net/http"
	"time"

	"github.com/adamwahyuh/coloroutfit/src/config"
	"github.com/adamwahyuh/coloroutfit/src/database"
	"github.com/adamwahyuh/coloroutfit/src/models"
	"github.com/adamwahyuh/coloroutfit/src/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorUnauthorized[any](),
		)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(config.GetEnv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorUnauthorized[any](),
		)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.ErrorUnauthorized[any](),
			)
			return
		}

		var user models.User
		database.DB.First(&user, claims["user_id"])

		if user.Id == 0 {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.ErrorUnauthorized[any](),
			)
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(
			http.StatusUnauthorized,
			response.ErrorUnauthorized[any](),
		)
		return
	}
}
