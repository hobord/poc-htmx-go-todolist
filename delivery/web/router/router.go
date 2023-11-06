package router

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

var (
	ParamsFromContext = httprouter.ParamsFromContext
)

type Group struct {
	router     *httprouter.Router
	prefix       string
	midlewares alice.Chain
}

func NewGroup(router *httprouter.Router, prefix string) *Group {
	return &Group{
		router:     router,
		prefix:       prefix,
		midlewares: alice.New(),
	}
}

func (rg *Group) Group(path string) *Group {
	group := NewGroup(rg.router, fmt.Sprintf("%s%s", rg.prefix, path))
	group.midlewares = rg.midlewares

	return group
}

func (rg *Group) WithMiddlewares(middleware ...alice.Constructor) *Group {
	rg.midlewares = rg.midlewares.Append(middleware...)

	return rg
}

func (rg *Group) Handler(method, path string, handler http.Handler) {
	rg.router.Handler(method, fmt.Sprintf("%s%s", rg.prefix, path), rg.midlewares.Then(handler))
}

func (rg *Group) GET(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodGet, path, handler)
}

func (rg *Group) POST(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodPost, path, handler)
}

func (rg *Group) PUT(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodPut, path, handler)
}

func (rg *Group) DELETE(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodDelete, path, handler)
}

func (rg *Group) PATCH(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodPatch, path, handler)
}

func (rg *Group) OPTIONS(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodOptions, path, handler)
}

func (rg *Group) HEAD(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodHead, path, handler)
}

func (rg *Group) CONNECT(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodConnect, path, handler)
}

func (rg *Group) TRACE(path string, handler http.HandlerFunc) {
	rg.Handler(http.MethodTrace, path, handler)
}
