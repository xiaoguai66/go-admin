package router

import (
	"admin-demo/api"
	"github.com/gin-gonic/gin"
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
		rgAuthUser.POST("", userApi.AddUser)
		rgAuthUser.GET("", userApi.GetUserList)
		rgAuthUser.GET("/:id", userApi.GetUserInfo)
	})
}
