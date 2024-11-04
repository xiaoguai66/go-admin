package api

import (
	"admin-demo/global"
	"admin-demo/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx                *gin.Context
	Request            any
	BuildParamsFromUri bool
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error
	//绑定请求上下文
	m.Ctx = option.Ctx

	//绑定请求数据
	if option.Request != nil {
		if option.BuildParamsFromUri {
			errResult = m.Ctx.ShouldBindUri(option.Request)
		} else {
			errResult = m.Ctx.ShouldBind(option.Request)
		}
		if errResult != nil {
			errResult = m.ParseValidateErrors(errResult, option.Request)
			m.AddError(errResult)
			m.Fail(ResponseJson{
				Msg: m.GetError().Error(),
			})
		}
	}

	return m
}

func (m *BaseApi) Ok(resp ResponseJson) {
	Ok(m.Ctx, resp)
}

func (m *BaseApi) Fail(resp ResponseJson) {
	Fail(m.Ctx, resp)
}

func (m *BaseApi) ServerError(resp ResponseJson) {
	ServerError(m.Ctx, resp)
}

func (m *BaseApi) AddError(err error) {
	m.Errors = utils.AppendError(m.Errors, err)
}

func (m *BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	//switch err.(type) {
	//case validator.ValidationErrors:
	//	fmt.Println("validator.ValidationErrors")
	//default:
	//	fmt.Println("未知Err")
	//}

	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}
	//通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, err := range errValidation {
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
