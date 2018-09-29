package router

const (
	absoluteHttpMethod = "*"
)

type IHandler interface {
	Handle(ctx IContext) bool
}

type HandlerFunc func(ctx IContext)

//обработчик, который фильтрует выполнение запроса по его методу
type FilteringHandler struct {
	handler    HandlerFunc
	httpMethod string
}

func NewFilteringHandler(h HandlerFunc, method string) *FilteringHandler {
	return &FilteringHandler{h, method}
}

//с bool Это временный костыль, который я позже исправлю
//лучше использовать мапу слайсов хэндлеров в роутере а не этот ГОвнокод
func (h *FilteringHandler) Handle(ctx IContext) bool{
	if h.httpMethod == absoluteHttpMethod || h.httpMethod == ctx.Method() {
		h.handler(ctx)
		return true
	}

	return false
}
