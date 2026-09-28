package routes

import "github.com/gin-gonic/gin"

func RegisterApiRoutes(route *gin.Engine) {
	api := route.Group("/api")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"message": "success"})
		})
	}
}
