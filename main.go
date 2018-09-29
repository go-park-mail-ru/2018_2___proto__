package main

import (
	"net/http"
	"proto-game-server/server"
)

func DefHandler(ctx server.IContext) {
	ctx.Write([]byte("IT'S ALIIIIIVEEEE!"))
	vars := ctx.UrlParams()

	ctx.Write([]byte(vars["id"]))
}

func main() {
	router := server.NewRouter()
	router.AddHandler("/user/{id}", DefHandler)

	http.ListenAndServe(":8080", router)
}
