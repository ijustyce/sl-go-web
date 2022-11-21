package SlFiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
	"net/http"
)

type FiberContext struct {
	context *fiber.Ctx
}

func (context *FiberContext) Response(code int, obj any) any {
	return context.context.JSON(obj)
}

func (context *FiberContext) QueryParam(key string, defaultValue string) string {
	return context.context.Query(key, defaultValue)
}

func (context *FiberContext) PathParam(key string, defaultValue string) string {
	return context.context.Params(key, defaultValue)
}

func (context *FiberContext) GetHeader(key string, defaultValue string) string {
	return context.context.Get(key, defaultValue)
}

func (context *FiberContext) BindBody(object any) error {
	return context.context.BodyParser(object)
}

func (context *FiberContext) BindQueryParams(object any) error {
	return context.context.QueryParser(object)
}

func (context *FiberContext) BindPathParams(object any) error {
	return context.context.ParamsParser(object)
}

func (context *FiberContext) BindHeaders(object any) error {
	return context.context.ReqHeaderParser(object)
}

func (context *FiberContext) SetHeader(key string, value string) {
	context.context.Set(key, value)
}

func (context *FiberContext) GetCookie(key string, defaultValue string) string {
	return context.context.Cookies(key, defaultValue)
}

func (context *FiberContext) SetCookie(cookie goweb.Cookie) {
	fiberCookie := fiber.Cookie{}
	fiberCookie.Name = cookie.Name
	fiberCookie.Value = cookie.Value
	fiberCookie.Domain = cookie.Domain
	fiberCookie.Expires = cookie.Expires
	fiberCookie.Path = cookie.Path
	fiberCookie.HTTPOnly = cookie.HTTPOnly
	fiberCookie.Secure = cookie.Secure
	fiberCookie.SessionOnly = cookie.SessionOnly
	fiberCookie.SameSite = toSameSite(cookie.SameSite)
	fiberCookie.MaxAge = cookie.MaxAge
	context.context.Cookie(&fiberCookie)
}

func (context *FiberContext) GetRequestIp() string {
	return context.context.IP()
}

func toSameSite(SameSite http.SameSite) string {
	switch SameSite {
	case http.SameSiteLaxMode:
	case http.SameSiteDefaultMode:
		return "Lax"
	case http.SameSiteNoneMode:
		return "None"
	case http.SameSiteStrictMode:
		return "Strict"
	}
	return "Lax"
}
