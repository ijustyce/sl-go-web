package goweb

type GoWeb interface {
	Get(path string, handler HandlerFunc)
	Post(path string, handler HandlerFunc)
	Delete(path string, handler HandlerFunc)
	Put(path string, handler HandlerFunc)
	Patch(path string, handler HandlerFunc)
	Head(path string, handler HandlerFunc)
	Options(path string, handler HandlerFunc)
	Start(port int) error
}

type SlContext interface {
	Response(code int, object any) any
	QueryParam(key string, defaultValue string) string
	PathParam(key string, defaultValue string) string
	BindBody(object any) error
	BindQueryParams(object any) error
	BindPathParams(object any) error
}

type Map map[string]any

type HandlerFunc func(SlContext)
