package main

import (
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb/SlGin"
)

func main() {
	goWeb := SlGin.Create()
	goWeb.Get("/api/:name", func(context goweb.SlContext) {
		name := context.PathParam("name2", "world")
		context.JSON(200, goweb.Map{
			"message": "hello " + name,
		})
	})

	goWeb.Start()
}
