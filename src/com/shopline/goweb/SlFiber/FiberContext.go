package SlFiber

import (
	"github.com/gofiber/fiber/v2"
)

type FiberContext struct {
	context *fiber.Ctx
}

func (context *FiberContext) JSON(code int, obj any) any {
	return context.context.JSON(obj)
}

func (context *FiberContext) QueryParam(key string, defaultValue string) string {
	return context.context.Query(key, defaultValue)
}

func (context *FiberContext) PathParam(key string, defaultValue string) string {
	return context.context.Params(key, defaultValue)
}
