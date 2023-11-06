package router

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

var (
	ParamsFromContext = httprouter.ParamsFromContext
)

type Group struct {
	router     *httprouter.Router
	prefix     string
	midlewares alice.Chain
}

func NewGroup(router *httprouter.Router, prefix string) *Group {
	return &Group{
		router:     router,
		prefix:     prefix,
		midlewares: alice.New(),
	}
}

func (g *Group) Group(path string) *Group {
	group := NewGroup(g.router, fmt.Sprintf("%s%s", g.prefix, path))
	group.midlewares = g.midlewares

	return group
}

func (g *Group) WithMiddlewares(middleware ...alice.Constructor) *Group {
	g.midlewares = g.midlewares.Append(middleware...)

	return g
}

func (g *Group) Handler(method, path string, handler http.Handler) {
	slog.Info("Register handler", "method", method, "path", fmt.Sprintf("%s%s", g.prefix, path))
	g.router.Handler(method, fmt.Sprintf("%s%s", g.prefix, path), g.midlewares.Then(handler))
}

func (g *Group) GET(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodGet, path, handler)
}

func (g *Group) POST(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodPost, path, handler)
}

func (g *Group) PUT(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodPut, path, handler)
}

func (g *Group) DELETE(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodDelete, path, handler)
}

func (g *Group) PATCH(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodPatch, path, handler)
}

func (g *Group) OPTIONS(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodOptions, path, handler)
}

func (g *Group) HEAD(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodHead, path, handler)
}

func (g *Group) CONNECT(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodConnect, path, handler)
}

func (g *Group) TRACE(path string, handler http.HandlerFunc) {
	g.Handler(http.MethodTrace, path, handler)
}
