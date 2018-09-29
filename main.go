package main

import (
	"net/http"
	"proto-game-server/router"
)

func DefHandler(ctxrouter.IContext) {
	ctx.Write([]byte("IT'S ALIIIIIVEEEE!"))
	vars := ctx.UrlParams()

	ctx.Write([]byte(vars["id"]))
}

func HelloHandler(ctxrouter.IContext) {
	ctx.Write([]byte("HELLO"))
}

func main() {
	router := router.NewRouter()
	router.AddHandler("/user/{id}", DefHandler)
	router.AddHandler("/test/asd/asd", HelloHandler)

	http.ListenAndServe(":8080", router)
}
