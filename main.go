package main

import (
	"net/http"
	"proto-game-server/router"
)

func main() {
	cfg, err := LoadConfigs("./data/cfg.json")
	if err != nil {
		panic(err)
	}

	router := router.NewRouter()
	apiHandler := NewApiHandler(cfg)

	//урлы должны быть отсортированы по длине урла по убыванию потом жобавлю это программно
	router.AddHandlerGet("/user/{slug}", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.GetUser)))
	router.AddHandlerGet("/leaders/{offset}/{limit}", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.GetUser)))
	router.AddHandlerGet("/test", apiHandler.AddCookie)

	router.AddHandlerPost("/signup", apiHandler.CorsEnableMiddleware(apiHandler.AddUser))
	router.AddHandlerPost("/signin", apiHandler.CorsEnableMiddleware(apiHandler.Authorize))

	router.AddHandlerDelete("/user", apiHandler.CorsEnableMiddleware(apiHandler.DeleteUser))
	router.AddHandlerPut("/user", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.UpdateUser)))
	router.AddHandlerOptions("/", apiHandler.CorsSetup)

	//запускаем сервер
	http.ListenAndServe(":8080", router)
}
