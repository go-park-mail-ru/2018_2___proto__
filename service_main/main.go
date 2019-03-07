package main

import (
	"net/http"
	"os"

	"proto-game-server/api"
	"proto-game-server/router"

	"github.com/op/go-logging"
	"github.com/prometheus/client_golang/prometheus"

	_ "net/http/pprof"
)

var hitCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "hits_total",
	Help: "Number of hits successfully processed.",
})

func CreateLogger() router.ILogger {
	format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)

	log := logging.MustGetLogger("logger")
	logging.SetBackend(formatter)

	return log
}

func Pprof() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
}

func main() {
	cfg, err := api.LoadConfigs("./data/cfg.json")
	if err != nil {
		panic(err)
	}

	if cfg.PprofEnabled {
		Pprof()
	}

	logger := CreateLogger()
	apiRouter := router.NewRouter(logger)
	nh := NewNetworkHandler(cfg, logger)

	prometheus.MustRegister(hitCounter)

	// TODO: урлы должны быть отсортированы по длине урла по убыванию потом добавлю это программно
	apiRouter.AddHandlerGet("/user/{slug}", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.GetUser)))
	apiRouter.AddHandlerGet("/user", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.Profile)))
	apiRouter.AddHandlerGet("/leaders/{offset}/{limit}", nh.CorsEnableMiddleware(nh.GetLeaders))
	apiRouter.AddHandlerGet("/session", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.GetSession)))
	apiRouter.AddHandlerGet("/game", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.ConnectPlayer)))
	// apiRouter.AddHandlerGet("/static/{file}", nh.CorsEnableMiddleware(nh.GetStatic))

	apiRouter.AddHandlerPost("/signup", nh.CorsEnableMiddleware(nh.AddUser))
	apiRouter.AddHandlerPost("/signin", nh.CorsEnableMiddleware(nh.Authorize))

	apiRouter.AddHandlerPut("/user", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.UpdateUser)))
	apiRouter.AddHandlerDelete("/user", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.DeleteUser)))
	apiRouter.AddHandlerDelete("/logout", nh.CorsEnableMiddleware(nh.AuthMiddleware(nh.Logout)))
	apiRouter.AddHandlerOptions("/", nh.CorsSetup)

	//урлы для тестирования
	//для нагрузочного тестирования
	apiRouter.AddHandlerGet("/loaderio-3b73ee37ac50f8785f6e274aba668913.txt", nh.verifyDomain)
	apiRouter.AddHandler("/metrics", nh.Metrics)

	// http.Handle("/metrics", promhttp.Handler())
	//запускаем сервер
	if cfg.UseHTTPS {
		err = http.ListenAndServeTLS(cfg.Port, "fullchain.pem", "privkey.pem", apiRouter)
	} else {
		err = http.ListenAndServe(cfg.Port, apiRouter)
	}

	logger.Critical(err)
}
