package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	ParamsFromContext = httprouter.ParamsFromContext
)

type RouteGroup struct {
	router     *httprouter.Router
	path       string
	midlewares []http.HandlerFunc
	groups     []*RouteGroup
	handlers   []func()
}

func NewRouteGroup(router *httprouter.Router, path string, midlewares ...http.HandlerFunc) *RouteGroup {
	return &RouteGroup{
		router: router,
	}
}

func (rg *RouteGroup) GET(path string, handle httprouter.Handle) {
	rg.router.GET(path, handle)
}
