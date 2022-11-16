package goweb

type GoWeb interface {
	Get(path string, handler HandlerFunc)

	Start()
}

type SlContext interface {
	JSON(code int, object any) any
	QueryParam(key string, defaultValue string) string
	PathParam(key string, defaultValue string) string
}

type Map map[string]any

type HandlerFunc func(SlContext)
