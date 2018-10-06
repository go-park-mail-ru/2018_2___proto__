package main

import (
	"log"
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
	router.AddHandlerGet("/user", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.Profile)))
	router.AddHandlerGet("/leaders/{offset}/{limit}", apiHandler.CorsEnableMiddleware(apiHandler.GetLeaders))
	router.AddHandlerGet("/test", apiHandler.AddCookie)

	router.AddHandlerPost("/signup", apiHandler.CorsEnableMiddleware(apiHandler.AddUser))
	router.AddHandlerPost("/signin", apiHandler.CorsEnableMiddleware(apiHandler.Authorize))

	router.AddHandlerDelete("/user", apiHandler.CorsEnableMiddleware(apiHandler.DeleteUser))
	router.AddHandlerPut("/user", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.UpdateUser)))
	router.AddHandlerOptions("/", apiHandler.CorsSetup)

	//запускаем сервер
	if cfg.UseHTTPS {
		err = http.ListenAndServeTLS(cfg.Port, "fullchain.pem", "privkey.pem", router)
	} else {
		err = http.ListenAndServe(cfg.Port, router)
	}

	log.Fatal(err)
}
