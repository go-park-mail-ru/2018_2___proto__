package router

import (
	"log"
	"net/http"
)

var supportedMethods = []string{
	http.MethodGet,
	http.MethodDelete,
	http.MethodPost,
	http.MethodPut,
	http.MethodOptions,
}

type Router struct {
	routes map[string][]*Route
}

func NewRouter() *Router {
	routes := make(map[string][]*Route)
	for _, val := range supportedMethods {
		routes[val] = make([]*Route, 0)
	}

	return &Router{routes}
}

func (r *Router) addHandlerOnMethos(urlPattern string, h HandlerFunc, httpMethod string) {
	route, err := NewDefaultRoute(urlPattern, h)
	if err != nil {
		panic(err)
	}

	_, ok := r.routes[httpMethod]
	if ok {
		r.routes[httpMethod] = append(r.routes[httpMethod], route)
	}
}

func (r *Router) AddHandler(urlPattern string, h HandlerFunc) {
	route, err := NewDefaultRoute(urlPattern, h)
	if err != nil {
		panic(err)
	}

	for _, val := range supportedMethods {
		r.routes[val] = append(r.routes[val], route)
	}
}

func (r *Router) AddHandlerGet(urlPattern string, h HandlerFunc) {
	r.addHandlerOnMethos(urlPattern, h, http.MethodGet)
}

func (r *Router) AddHandlerPost(urlPattern string, h HandlerFunc) {
	r.addHandlerOnMethos(urlPattern, h, http.MethodPost)
}

func (r *Router) AddHandlerDelete(urlPattern string, h HandlerFunc) {
	r.addHandlerOnMethos(urlPattern, h, http.MethodDelete)
}

func (r *Router) AddHandlerPut(urlPattern string, h HandlerFunc) {
	r.addHandlerOnMethos(urlPattern, h, http.MethodPut)
}

func (r *Router) AddHandlerOptions(urlPattern string, h HandlerFunc) {
	r.addHandlerOnMethos(urlPattern, h, http.MethodOptions)
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	log.Printf("%v: %v", req.Method, req.RequestURI)
	ctx := NewContext(writer, req)

	routes, ok := r.routes[req.Method]
	if ok {
		for _, route := range routes {
			if route.Match(ctx) {
				route.Handle(ctx)
				return
			}
		}
	}

	//если не один хэндлер не отработал, то возвращаем 404
	ctx.StatusCode(http.StatusNotFound)
}
