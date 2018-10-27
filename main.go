package main

import (
	"net/http"
	"os"
	"proto-game-server/router"

	"github.com/op/go-logging"
)

func CreateLogger() router.ILogger {
	format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)

	log := logging.MustGetLogger("logger")
	logging.SetBackend(formatter)

	return log
}

func main() {
	cfg, err := LoadConfigs("./data/cfg.json")
	if err != nil {
		panic(err)
	}

	logger := CreateLogger()
	apiRouter := router.NewRouter(logger)
	apiHandler := NewApiHandler(cfg)

	//урлы должны быть отсортированы по длине урла по убыванию потом жобавлю это программно
	apiRouter.AddHandlerGet("/user/{slug}", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.GetUser)))
	apiRouter.AddHandlerGet("/user", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.Profile)))
	apiRouter.AddHandlerGet("/leaders/{offset}/{limit}", apiHandler.CorsEnableMiddleware(apiHandler.GetLeaders))
	apiRouter.AddHandlerGet("/session", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.GetSession)))
	apiRouter.AddHandlerGet("/static/{file}", apiHandler.CorsEnableMiddleware(apiHandler.GetStatic))
	apiRouter.AddHandlerGet("/test", apiHandler.Test)

	apiRouter.AddHandlerPost("/signup", apiHandler.CorsEnableMiddleware(apiHandler.AddUser))
	apiRouter.AddHandlerPost("/signin", apiHandler.CorsEnableMiddleware(apiHandler.Authorize))

	apiRouter.AddHandlerPut("/user", apiHandler.CorsEnableMiddleware(apiHandler.AuthMiddleware(apiHandler.UpdateUser)))
	apiRouter.AddHandlerDelete("/user", apiHandler.CorsEnableMiddleware(apiHandler.DeleteUser))
	apiRouter.AddHandlerOptions("/", apiHandler.CorsSetup)

	// этот путь необходим для проведения нагрузочного тестирования
	apiRouter.AddHandlerGet("/loaderio-3b73ee37ac50f8785f6e274aba668913.txt", apiHandler.verifyDomain)

	//запускаем сервер
	if cfg.UseHTTPS {
		err = http.ListenAndServeTLS(cfg.Port, "fullchain.pem", "privkey.pem", apiRouter)
	} else {
		err = http.ListenAndServe(cfg.Port, apiRouter)
	}

	logger.Critical(err)
}
