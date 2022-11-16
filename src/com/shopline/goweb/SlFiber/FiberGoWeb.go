package SlFiber

import "sl-go/src/com/shopline/goweb"

import (
	"github.com/gofiber/fiber/v2"
)

type SlFiber struct {
	app *fiber.App
}

func Create() *SlFiber {
	slFiber := &SlFiber{}
	slFiber.app = fiber.New(fiber.Config{
		Prefork:                  false,
		CaseSensitive:            true,
		StrictRouting:            true,
		DisableDefaultDate:       false,
		DisableHeaderNormalizing: true,
	})
	return slFiber
}

func (slFiber *SlFiber) Get(path string, handler goweb.HandlerFunc) {
	slFiber.app.Get(path, func(ctx *fiber.Ctx) error {
		fiberContext := &FiberContext{}
		fiberContext.context = ctx
		handler(fiberContext)
		return nil
	})
}

func (slFiber *SlFiber) Start() {
	slFiber.app.Listen(":8080")
}
