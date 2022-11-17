package SlGin

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	context *gin.Context
}

func (context *GinContext) Response(code int, obj any) any {
	context.context.JSON(code, obj)
	return nil
}

func (context *GinContext) QueryParam(key string, defaultValue string) string {
	return context.context.DefaultQuery(key, defaultValue)
}

func (context *GinContext) PathParam(key string, defaultValue string) string {
	value := context.context.Param(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (context *GinContext) BindBody(object any) error {
	return context.context.ShouldBindJSON(object)
}

func (context *GinContext) BindQueryParams(object any) error {
	return context.context.ShouldBindQuery(object)
}

func (context *GinContext) BindPathParams(object any) error {
	return context.context.ShouldBindUri(object)
}
