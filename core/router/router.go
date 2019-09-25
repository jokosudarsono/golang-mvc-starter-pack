package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router              *mux.Router
	ComposedMiddlewares []func(http.HandlerFunc) http.HandlerFunc
}

func (rt *Router) Route(method string, endpoint string, handler http.HandlerFunc) {
	rt.Router.Handle(endpoint, wrapper(handler, rt.ComposedMiddlewares)).Methods(method)
}

func (rt *Router) Default(handler http.HandlerFunc) {
	rt.Router.PathPrefix("/").HandlerFunc(handler)
}

func (rt *Router) GET(endpoint string, handler http.HandlerFunc) *Router {
	rt.Route("GET", endpoint, handler)
	return rt
}

func (rt *Router) POST(endpoint string, handler http.HandlerFunc) *Router {
	rt.Route("POST", endpoint, handler)
	return rt
}

func (rt *Router) PUT(endpoint string, handler http.HandlerFunc) *Router {
	rt.Route("PUT", endpoint, handler)
	return rt
}

func (rt *Router) PATCH(endpoint string, handler http.HandlerFunc) *Router {
	rt.Route("PATCH", endpoint, handler)
	return rt
}

func (rt *Router) DELETE(endpoint string, handler http.HandlerFunc) *Router {
	rt.Route("DELETE", endpoint, handler)
	return rt
}

func (rt *Router) Middlewares(
	mws ...func(http.HandlerFunc) http.HandlerFunc,
) *Router {
	rt.ComposedMiddlewares = mws
	return rt
}

func wrapper(
	handler http.HandlerFunc,
	mws []func(http.HandlerFunc) http.HandlerFunc,
) http.HandlerFunc {
	for _, middleware := range mws {
		handler = middleware(handler)
	}

	return handler
}
