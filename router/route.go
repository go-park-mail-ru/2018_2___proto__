package router

type Route struct {
	matcher IApiUrlMatcher
	parser  IApiUrlParser
	handler IHandler
}

func NewDefaultRoute(url string, handler HandlerFunc, method string) (*Route, error) {
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
