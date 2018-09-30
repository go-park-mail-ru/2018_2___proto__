package router

type HandlerFunc func(ctx IContext)

type Route struct {
	matcher IApiUrlMatcher
	parser  IApiUrlParser
	handler HandlerFunc
}

func NewDefaultRoute(url string, handler HandlerFunc) (*Route, error) {
	route := new(Route)

	matcher, err := NewApiUrlParser(url)
	if err != nil {
		return nil, err
	}

	route.parser = matcher
	route.matcher = matcher
	route.handler = handler
	return route, nil
}

func (r *Route) Match(ctx IContext) bool {
	return r.matcher.Match(ctx.RequestURI())
}

func (r *Route) Handle(ctx IContext) {
	ctx.SetApiParser(r.parser)
	r.handler(ctx)
}
