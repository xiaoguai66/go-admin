package router

import (
	"admin-demo/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRouters() {
	RegisterRouter(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublic.POST("login", userApi.Login)

		rgAuthUser := rgAuth.Group("user").Use(func() gin.HandlerFunc {
			return func(context *gin.Context) {
				//context.AbortWithStatusJSON(200, gin.H{
				//	"Msg": "Login Middleware",
				//})
			}
		}())
		rgAuthUser.GET("", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"data": []map[string]any{
					{"id": 1, "name": "zs"},
					{"id": 2, "name": "ls"},
				},
			})
		})
		rgAuthUser.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "zs",
			})
		})
	})
}
