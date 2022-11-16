package SlGin

import (
	"github.com/gin-gonic/gin"
	"sl-go/src/com/shopline/goweb"
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

func (slGin *SlGin) Start() {
	slGin.gin.Run()
}

func (slGin *SlGin) Get(path string, handler goweb.HandlerFunc) {
	slGin.gin.GET(path, func(context *gin.Context) {
		ginContext := &GinContext{}
		ginContext.context = context
		handler(ginContext)
	})
}
