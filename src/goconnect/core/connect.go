package core

import (
	"log"
	"net/http"
)

var (
	NOOP = func() {}
)

type Connect struct {
	middlewares []*Middleware
}

type MiddlewareFunc interface {
	Name() string
	ServeHTTP(res http.ResponseWriter, req *http.Request, next NextMiddleware)
}

type NextMiddleware func()

type Middleware struct {
	Handler MiddlewareFunc
}

func NewConnect() (*Connect, error) {
	return &Connect{}, nil
}

func (connect *Connect) Use(fn MiddlewareFunc) *Connect {
	log.Printf("%s %d\n", fn.Name(), len(connect.middlewares))
	connect.middlewares = append(connect.middlewares, &Middleware{Handler: fn})
	return connect
}

func (connect *Connect) Length() int {
	return len(connect.middlewares)
}

func (connect *Connect) MakeNext(res http.ResponseWriter, req *http.Request, index int) func() {
	length := len(connect.middlewares)
	if index >= length {
		return NOOP
	} else {
		return func() {
			middleware := connect.middlewares[index]
			log.Printf("adding handler %s", middleware.Handler.Name())
			next := connect.MakeNext(res, req, index+1)
			middleware.Handler.ServeHTTP(res, req, next)
		}
	}
}

func (connect *Connect) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	i := 0
	length := len(connect.middlewares)
	if length < 1 {
		return
	}
	handler := connect.middlewares[i]
	next := connect.MakeNext(res, req, i+1)
	handler.Handler.ServeHTTP(res, req, next)
}
