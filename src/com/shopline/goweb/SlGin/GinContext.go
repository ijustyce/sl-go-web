package SlGin

import (
	"github.com/gin-gonic/gin"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
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

func (context *GinContext) GetHeader(key string, defaultValue string) string {
	value := context.context.GetHeader(key)
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

func (context *GinContext) BindHeaders(object any) error {
	return context.context.ShouldBindHeader(object)
}

func (context *GinContext) SetHeader(key string, value string) {
	context.context.Header(key, value)
}

func (context *GinContext) GetCookie(key string, defaultValue string) string {
	value, err := context.context.Cookie(key)
	if err != nil {
		return defaultValue
	}
	return value
}

func (context *GinContext) SetCookie(cookie goweb.Cookie) {
	context.context.SetSameSite(cookie.SameSite)
	context.context.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path,
		cookie.Domain, cookie.Secure, cookie.HTTPOnly)
}

func (context *GinContext) GetRequestIp() string {
	return context.context.ClientIP()
}
