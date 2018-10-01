package main

import (
	"encoding/json"
	"log"
	"net/http"
	"proto-game-server/api"
	m "proto-game-server/models"
	"proto-game-server/router"
)

const (
	cookieSessionIdName = "sessionId"
)

//посредник между сетью и логикой апи
type ApiHandler struct {
	apiService      *api.ApiService
	corsAllowedHost string
}

//избавиться от хардкода коннекта к бд
func NewApiHandler(settings *ServerConfig) *ApiHandler {
	service, err := api.NewApiService(settings.DbConnector, settings.DbConnectionString)
	if err != nil {
		panic(err)
	}

	return &ApiHandler{
		apiService:      service,
		corsAllowedHost: settings.CorsAllowedHost,
	}
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

//регистрация
//обязательно нужно реализовать
func (h *ApiHandler) AddUser(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	//можно потом добавить валидацию, но не сейчас

	//передаем уюзера из тела запроса в хранилище юзеров на регистрацию
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

//миддлварь для аутентификации
//обязательно нужно реализовать
func (h *ApiHandler) AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		//тут должно быть получение id сессии из кукисов
		//попытка найти сессию в хранилище сессий и вызов след обработчика если все норм
		sessionCookie, err := ctx.GetCookie(cookieSessionIdName)
		if err != nil {
			WriteResponse(&api.ApiResponse{Code: http.StatusInternalServerError,
				Response: "ошибка поиска сессии в куках"}, ctx)
			return
		}

		//поиск сессии по ИД в хранилище
		_, isSessionExists := h.apiService.Sessions.GetById(sessionCookie.Value)

		if !isSessionExists {
			WriteResponse(&api.ApiResponse{http.StatusUnauthorized, "вы не авторизованы"}, ctx)
			return
		}

		next(ctx)
	}
}

//настройка cors'a
func (h *ApiHandler) CorsEnableMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		ctx.Header("Access-Control-Allow-Origin", h.corsAllowedHost)
		ctx.Header("Access-Control-Allow-Credentials", "true")

		next(ctx)
	}
}

//обработчик регистрации
//обязательно нужно реализовать
func (h *ApiHandler) Authorize(ctx router.IContext) {
	user := new(m.User)
	ctx.ReadJSON(user)

	//тут должна быть авторизация и выдача ид сессии в куки
	//хранилище создают сессию и возвращает нам ид сессии, который записывам в куки
	sessionId, ok, err := h.apiService.Sessions.Create(user)
	if !ok {
		WriteResponse(&api.ApiResponse{Code: 401, Response: &m.Error{Code: 401, Message: err}}, ctx)
		return
	}

	//записываем ид сессии в куки
	//при каждом запросе, требующем аутнетификацию, будет брвться данная кука и искаться в хранилище
	ctx.SetCookie(&http.Cookie{Name: cookieSessionIdName, Value: sessionId})
	ctx.StatusCode(http.StatusOK)
}
