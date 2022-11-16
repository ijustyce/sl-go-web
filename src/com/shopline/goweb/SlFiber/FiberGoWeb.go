package SlFiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
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
