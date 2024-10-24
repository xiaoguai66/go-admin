package api

import (
	"admin-demo/service/dto"
	"admin-demo/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type UserApi struct {
}

func NewUserApi() *UserApi {
	return &UserApi{}
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
	err := ctx.ShouldBind(&iUserLoginDto)
	if err != nil {
		//switch err.(type) {
		//case validator.ValidationErrors:
		//	fmt.Println("validator.ValidationErrors")
		//default:
		//	fmt.Println("未知Err")
		//}
		Fail(ctx, ResponseJson{Msg: parseValidateErrors(err.(validator.ValidationErrors), &iUserLoginDto).Error()})
		return
	}
	Ok(ctx, ResponseJson{Data: iUserLoginDto})

}

func parseValidateErrors(errs validator.ValidationErrors, target any) error {
	var errResult error

	//通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, err := range errs {
		field, _ := fields.FieldByName(err.Field())
		errMessageTag := fmt.Sprintf("%s_err", err.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage := field.Tag.Get("message")
			if errMessage == "" {
				errMessage = err.Field() + " 参数错误 " + err.Tag()
			}
		}
		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}
	return errResult
}
