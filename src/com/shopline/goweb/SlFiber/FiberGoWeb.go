package SlFiber

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
)

type SlFiber struct {
	app *fiber.App
}

func Create(debugMode bool) *SlFiber {
	slFiber := &SlFiber{}
	slFiber.app = fiber.New(fiber.Config{
		Prefork:                  false,
		CaseSensitive:            true,
		StrictRouting:            true,
		DisableDefaultDate:       false,
		DisableHeaderNormalizing: false,
	})
	return slFiber
}

func fiberHandler(handler goweb.HandlerFunc) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fiberContext := &FiberContext{}
		fiberContext.context = ctx
		handler(fiberContext)
		return nil
	}
}

func (slFiber *SlFiber) Get(path string, handler goweb.HandlerFunc) {
	slFiber.app.Get(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Post(path string, handler goweb.HandlerFunc) {
	slFiber.app.Post(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Delete(path string, handler goweb.HandlerFunc) {
	slFiber.app.Delete(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Put(path string, handler goweb.HandlerFunc) {
	slFiber.app.Put(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Patch(path string, handler goweb.HandlerFunc) {
	slFiber.app.Patch(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Head(path string, handler goweb.HandlerFunc) {
	slFiber.app.Head(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Options(path string, handler goweb.HandlerFunc) {
	slFiber.app.Options(path, fiberHandler(handler))
}

func (slFiber *SlFiber) Start(port int) error {
	addr := fmt.Sprintf("%s%d", ":", port)
	return slFiber.app.Listen(addr)
}
