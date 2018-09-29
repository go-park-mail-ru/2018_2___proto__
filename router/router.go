package router

import (
	"net/http"
)

type Router struct {
	routes []*Route
}

func NewRouter() *Router {
	return &Router{make([]*Route, 0)}
}

func (r *Router) addHandlerFull(urlPattern string, h HandlerFunc, httpMethod string) {
	route, err := NewDefaultRoute(urlPattern, h, httpMethod)
	if err != nil {
		panic(err)
	}

	r.routes = append(r.routes, route)
}

func (r *Router) AddHandler(urlPattern string, h HandlerFunc) {
	r.addHandlerFull(urlPattern, h, absoluteHttpMethod)
}

func (r *Router) AddHandlerGet(urlPattern string, h HandlerFunc) {
	r.addHandlerFull(urlPattern, h, http.MethodGet)
}

func (r *Router) AddHandlerPost(urlPattern string, h HandlerFunc) {
	r.addHandlerFull(urlPattern, h, http.MethodPost)
}

func (r *Router) AddHandlerDelete(urlPattern string, h HandlerFunc) {
	r.addHandlerFull(urlPattern, h, http.MethodDelete)
}

func (r *Router) AddHandlerPut(urlPattern string, h HandlerFunc) {
	r.addHandlerFull(urlPattern, h, http.MethodPut)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	ctx := NewContext(writer, req)

	for _, val := range r.routes {
		if val.TryHandle(ctx) {
			return
		}
	}

	//если не один хэндлер не отработал, то возвращаем 404
	ctx.StatusCode(http.StatusNotFound)
}
