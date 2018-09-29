package main

import (
	"net/http"
	"encoding/json"
	"log"
	"proto-game-server/api"
	m "proto-game-server/models"
	"proto-game-server/router"
)

//посредник между сетью и логикой апи
type ApiHandler struct {
	apiService *api.ApiService
}

func NewApiHandler() *ApiHandler {
	return &ApiHandler{api.NewApiService()}
}

func WriteResponse(response *api.ApiResponse, ctx router.IContext) {
	data, err := json.Marshal(response.Response)
	if err != nil {
		log.Fatalln(err)
	}

	ctx.ContentType("application/json")
	ctx.StatusCode(response.Code)
	ctx.Write(data)
}

func (h *ApiHandler) AddUser(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	WriteResponse(h.apiService.Users.Add(user), ctx)
}

func (h *ApiHandler) DeleteUser(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	WriteResponse(h.apiService.Users.Remove(user), ctx)
}

func (h *ApiHandler) UpdateUser(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	WriteResponse(h.apiService.Users.Update(user), ctx)
}

func (h *ApiHandler) GetUser(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	params := ctx.UrlParams()
	WriteResponse(h.apiService.Users.Get(params["slug"]), ctx)
}

func (h *ApiHandler) AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		//тут должно быть получение id сессии из кукисов
		//попытка найти сессию в хранилище сессий и вызов след обработчика если все норм
		isSessionExists := false

		if !isSessionExists {
			WriteResponse(&api.ApiResponse{401, "вы не авторизованы"}, ctx)
			return
		}

		next(ctx)
	}
}

func (h *ApiHandler) Authorize(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	//тут должна быть авторизация и выдача ид сессии в куки
	sessionId := h.apiService.Sessions.Create(user)
	ctx.SetCookie(&http.Cookie{Name: "sessionId",Value:sessionId})
}