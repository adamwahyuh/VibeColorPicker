package routes

import (
	authcontroller "github.com/adamwahyuh/coloroutfit/src/controllers/AuthController"
	usercontroller "github.com/adamwahyuh/coloroutfit/src/controllers/UserController"
	"github.com/adamwahyuh/coloroutfit/src/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(route *gin.Engine) {
	api := route.Group("/api")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})

		auth := api.Group("/authentication")
		{
			auth.POST("/register", usercontroller.StoreUser)
			auth.POST("/login", authcontroller.Login)
			auth.GET("/check", middleware.RequireAuth, authcontroller.CheckLogin)
			auth.POST("/logout", middleware.RequireAuth, authcontroller.Logout)
		}
	}
}
