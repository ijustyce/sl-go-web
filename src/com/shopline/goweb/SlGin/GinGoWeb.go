package SlGin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
)

type SlGin struct {
	gin *gin.Engine
}

func Create() *SlGin {
	slGin := &SlGin{}
	gin.SetMode(gin.ReleaseMode)
	slGin.gin = gin.New()
	slGin.gin.Use(gin.Recovery())
	return slGin
}

func ginHandler(handler goweb.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		ginContext := &GinContext{}
		ginContext.context = context
		handler(ginContext)
	}
}

func (slGin *SlGin) Get(path string, handler goweb.HandlerFunc) {
	slGin.gin.GET(path, ginHandler(handler))
}

func (slGin *SlGin) Post(path string, handler goweb.HandlerFunc) {
	slGin.gin.POST(path, ginHandler(handler))
}

func (slGin *SlGin) Delete(path string, handler goweb.HandlerFunc) {
	slGin.gin.DELETE(path, ginHandler(handler))
}

func (slGin *SlGin) Put(path string, handler goweb.HandlerFunc) {
	slGin.gin.PUT(path, ginHandler(handler))
}

func (slGin *SlGin) Patch(path string, handler goweb.HandlerFunc) {
	slGin.gin.PATCH(path, ginHandler(handler))
}

func (slGin *SlGin) Head(path string, handler goweb.HandlerFunc) {
	slGin.gin.HEAD(path, ginHandler(handler))
}

func (slGin *SlGin) Options(path string, handler goweb.HandlerFunc) {
	slGin.gin.OPTIONS(path, ginHandler(handler))
}

func (slGin *SlGin) Start(port int) error {
	addr := fmt.Sprintf("%s%d", ":", port)
	return slGin.gin.Run(addr)
}
