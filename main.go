package main

import (
	"fmt"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb"
	"github.com/ijustyce/sl-go-web/src/com/shopline/goweb/SlFiber"
)

type PathParams struct {
	Name string `uri:"name" params:"name"`
	Age  string `uri:"age" params:"age"`
}

type QueryParams struct {
	Name string `form:"name" query:"name"`
	Age  string `form:"age" query:"age"`
}

type Body struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type Header struct {
	ContentType string `header:"Content-type" reqHeader:"Content-type"`
	UserAgent   string `header:"User-Agent" reqHeader:"User-Agent"`
	Host        string `header:"Host" reqHeader:"Host"`
}

func main() {
	goWeb := SlFiber.Create()
	goWeb.Get("/api/:name/:age", func(context goweb.SlContext) {
		var pathParams PathParams
		err := context.BindPathParams(&pathParams)
		if err != nil {
			fmt.Println(err)
		}

		var queryParams QueryParams
		err2 := context.BindQueryParams(&queryParams)
		if err2 != nil {
			fmt.Println(err2)
		}

		userId := context.GetCookie("userId", "")
		if userId == "" {
			fmt.Println("user not login")
			cookie := goweb.Cookie{Name: "userId", Value: "chun"}
			context.SetCookie(cookie)
		}

		context.SetHeader("trace-id", "your traceId")
		context.Response(200, goweb.Map{
			"message": "hello " + pathParams.Name,
		})
	})

	InitWev(goWeb)

	goWeb2 := SlFiber.Create()
	InitWev(goWeb2)

	err := goWeb.Start(8080)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitWev(web goweb.GoWeb) {
}
