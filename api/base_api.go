package api

import (
	"admin-demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

type BuildRequestOption struct {
	Ctx                *gin.Context
	DTO                any
	BuildParamsFromUri bool
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {

	return m
}

func (m *BaseApi) AddError(err error) {
	m.Errors = utils.AppendError(m.Errors, err)
}

func (m *BaseApi) GetError() error {
	return m.Errors
}
