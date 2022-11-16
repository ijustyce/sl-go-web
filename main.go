package main

import (
	"sl-go/src/com/shopline/goweb"
	"sl-go/src/com/shopline/goweb/SlGin"
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
