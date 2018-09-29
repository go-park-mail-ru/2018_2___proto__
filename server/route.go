package server

import (
	"net/http"
)

type Route struct {
	matcher IApiUrlMatcher
	parser IApiUrlParser
	handler IHandler
}

func newDefaultRoute(url string, handler HandlerFunc, method string) (*Route, error) {
	route := new(Route)

	matcher, err := NewApiUrlParser(url)
	if err != nil {
		return nil, err
	}

	route.parser = matcher
	route.matcher = matcher
	route.handler = NewFilteringHandler(handler, method)
	return route, nil
}

func NewDefaultRoute(url string, handler HandlerFunc) (*Route, error) {
	return newDefaultRoute(url, handler, absoluteHttpMethod)
}

func NewDefaultRouteGet(url string, handler HandlerFunc) (*Route, error) {
	return newDefaultRoute(url, handler, http.MethodGet)
}

func NewDefaultRoutePost(url string, handler HandlerFunc) (*Route, error) {
	return newDefaultRoute(url, handler, http.MethodPost)
}

func NewDefaultRoutePut(url string, handler HandlerFunc) (*Route, error) {
	return newDefaultRoute(url, handler, http.MethodPut)
}

func NewDefaultRouteDelete(url string, handler HandlerFunc) (*Route, error) {
	return newDefaultRoute(url, handler, http.MethodDelete)
}

//если дорога правильная, то возвращается true
func (r *Route) TryHandle(ctx IContext) bool {
	if !r.matcher.Match(ctx.RequestURI()) {
		return false
	}

	//не придумал, как это лучше организовать, поэтому парсер назначается тут
	//внутри обработчика можно вызывать парс 
	ctx.SetApiParser(r.parser)
	r.handler.Handle(ctx)
	return true
}