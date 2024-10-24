package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
}

func HttpResponse(ctx *gin.Context, resp ResponseJson, status int) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}
	ctx.AbortWithStatusJSON(status, resp)
}

func (resp ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(resp, ResponseJson{})
}

func buildStatus(resp ResponseJson, defaultStatus int) int {
	if 0 == resp.Status {
		return defaultStatus
	}
	return resp.Status
}

func Ok(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, resp, buildStatus(resp, http.StatusOK))
	//ctx.AbortWithStatusJSON(buildStatus(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, resp, buildStatus(resp, http.StatusBadRequest))
	//ctx.AbortWithStatusJSON(buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerError(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, resp, buildStatus(resp, http.StatusInternalServerError))
}
