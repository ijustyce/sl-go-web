package SlFiber

import (
	"github.com/gofiber/fiber/v2"
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

func (context *FiberContext) BindBody(object any) error {
	return context.context.BodyParser(object)
}

func (context *FiberContext) BindQueryParams(object any) error {
	return context.context.QueryParser(object)
}

func (context *FiberContext) BindPathParams(object any) error {
	return context.context.ParamsParser(object)
}
