package api

import (
	"admin-demo/service"
	"admin-demo/service/request"
	"admin-demo/utils"
	"github.com/gin-gonic/gin"
)

const (
	ERR_CODE_ADD_USER       = 10011
	ERR_CODE_GET_USER_BY_ID = 10012
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
	var iUserLoginRequest request.UserLoginRequest
	//err := ctx.ShouldBind(&iUserLoginRequest)
	//if err != nil {
	//
	//}
	if err := u.BuildRequest(BuildRequestOption{
		Ctx:     ctx,
		Request: &iUserLoginRequest,
	}).GetError(); err != nil {
		return
	}
	user, err := u.Service.Login(iUserLoginRequest)
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
	var iUserAddRequest request.UserAddRequest
	if err := u.BuildRequest(BuildRequestOption{
		Ctx:     ctx,
		Request: &iUserAddRequest,
	}).GetError(); err != nil {
		return
	}
	err := u.Service.AddUser(&iUserAddRequest)
	if err != nil {
		u.Fail(ResponseJson{
			Code: ERR_CODE_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	u.Ok(ResponseJson{
		Data: iUserAddRequest,
	})
}

func (u UserApi) GetUserList(ctx *gin.Context) {

}

func (u UserApi) GetUserInfo(ctx *gin.Context) {
	var idRequest request.CommonIDRequest
	if err := u.BuildRequest(BuildRequestOption{
		Ctx:                ctx,
		Request:            &idRequest,
		BuildParamsFromUri: true,
	}).GetError(); err != nil {
		return
	}
	userModel, err := u.Service.GetUserInfoById(&idRequest)
	if err != nil {
		u.Fail(ResponseJson{
			Code: ERR_CODE_GET_USER_BY_ID,
			Msg:  err.Error(),
		})
		return
	}
	u.Ok(ResponseJson{
		Data: userModel,
	})
}
