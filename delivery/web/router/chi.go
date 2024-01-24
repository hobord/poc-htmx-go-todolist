package router

import (
	"fmt"
	"log/slog"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

var (
	PathValue = func(r *http.Request, key string) string {
		return chi.URLParam(r, key)
	}
)

func NewRouter() chi.Router {
	return chi.NewRouter()
}

type Group struct {
	router     chi.Router
	prefix     string
	midlewares []func(http.Handler) http.Handler
}

func NewGroup(router chi.Router, prefix string) *Group {
	return &Group{
		router:     router,
		prefix:     prefix,
		midlewares: make([]func(http.Handler) http.Handler, 0),
	}
}

func (g *Group) Group(path string) *Group {
	group := NewGroup(g.router, fmt.Sprintf("%s%s", g.prefix, path))
	group.midlewares = g.midlewares

	return group
}

func (g *Group) WithMiddlewares(middleware ...func(http.Handler) http.Handler) *Group {
	g.midlewares = append(g.midlewares, middleware...)

	return g
}

func (g *Group) Handler(method, path string, handler http.Handler) {
	slog.Info("Register handler", "method", method, "path", fmt.Sprintf("%s%s", g.prefix, path))

	g.router.With(g.midlewares...).Method(method, fmt.Sprintf("%s%s", g.prefix, path), handler)
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
