package server

import (
	"net/http"
)

const (
	absoluteHttpMethod = "*"
)

type IHandler interface {
	Handle(ctx IContext)
}

type HandlerFunc func(ctx IContext)

//обработчик, который фильтрует выполнение запроса по его методу
type FilteringHandler struct {
	handler HandlerFunc
	httpMethod string
}

func NewFilteringHandler(h HandlerFunc, method string) *FilteringHandler {
	return &FilteringHandler{h, method}
}

func (h *FilteringHandler) Handle(ctx IContext) {
	if h.httpMethod == absoluteHttpMethod || h.httpMethod == ctx.Method() {
		h.handler(ctx)
	}

	ctx.Status(http.StatusNotFound)
}