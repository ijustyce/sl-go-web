package SlGin

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	context *gin.Context
}

func (context *GinContext) JSON(code int, obj any) any {
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
