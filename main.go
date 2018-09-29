package main

import (
	"net/http"
	"proto-game-server/router"
)

func main() {
	router := router.NewRouter()
	apiHandler := NewApiHandler()

	//урлы должны быть отсортированы по длине урла по убыванию потом жобавлю это программно
	router.AddHandlerGet("/user/{slug}", apiHandler.AuthMiddleware(apiHandler.GetUser))
	router.AddHandlerPost("/user", apiHandler.AddUser)
	router.AddHandlerDelete("/user", apiHandler.DeleteUser)
	router.AddHandlerPut("/user", apiHandler.AuthMiddleware(apiHandler.UpdateUser))

	http.ListenAndServe(":8080", router)
}
