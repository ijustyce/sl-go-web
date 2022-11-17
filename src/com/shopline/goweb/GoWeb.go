package goweb

import (
	"net/http"
	"time"
)

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
	GetHeader(key string, defaultValue string) string
	BindBody(object any) error
	BindQueryParams(object any) error
	BindPathParams(object any) error
	BindHeaders(object any) error
	SetHeader(key string, value string)
	GetCookie(key string, defaultValue string) string
	SetCookie(cookie Cookie)
}

type Cookie struct {
	Name        string
	Value       string
	Path        string
	Domain      string
	MaxAge      int
	Expires     time.Time
	Secure      bool
	HTTPOnly    bool
	SameSite    http.SameSite
	SessionOnly bool
}

type Map map[string]any

type HandlerFunc func(SlContext)
