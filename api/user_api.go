package api

import (
	"admin-demo/service/dto"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
}

func NewUserApi() *UserApi {
	return &UserApi{
		BaseApi: NewBaseApi(),
	}
}

// Login
// @Tags 用户管理
// @Summary 用户登录.
// @Description 用户登录1
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/login [post]
func (u UserApi) Login(ctx *gin.Context) {
	//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})
	//Ok(ctx, ResponseJson{
	//	Msg: "Login Success",
	//})
	var iUserLoginDto dto.UserLoginDto
	//err := ctx.ShouldBind(&iUserLoginDto)
	//if err != nil {
	//
	//}
	if err := u.BuildRequest(BuildRequestOption{
		Ctx: ctx,
		DTO: &iUserLoginDto,
	}).GetError(); err != nil {
		return
	}

	u.Ok(ResponseJson{Data: iUserLoginDto})

}
