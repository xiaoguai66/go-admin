package api

import (
	"admin-demo/service"
	"admin-demo/service/request"
	"admin-demo/utils"
	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER = 10011
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{
		BaseApi: NewBaseApi(),
		Service: service.NewUserService(),
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
	var iUserLoginDto request.UserLoginDto
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
	user, err := u.Service.Login(iUserLoginDto)
	if err != nil {
		u.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(int(user.ID), user.Name)
	if err != nil {
		u.Fail(ResponseJson{
			Msg: err.Error(),
		})
		return
	}

	u.Ok(ResponseJson{Data: gin.H{
		"token": token,
		"user":  user,
	}})

}

func (u UserApi) AddUser(ctx *gin.Context) {
	var iUserAddDto request.UserAddDto
	if err := u.BuildRequest(BuildRequestOption{
		Ctx: ctx,
		DTO: &iUserAddDto,
	}).GetError(); err != nil {
		return
	}
	err := u.Service.AddUser(&iUserAddDto)
	if err != nil {
		u.Fail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	u.Ok(ResponseJson{
		Data: iUserAddDto,
	})
}
