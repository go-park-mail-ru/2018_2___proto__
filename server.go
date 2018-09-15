package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

//temporary token for authentification
const (
	TmpToken                    = "yogsothoth42shubnigurath"
	AuthorizationFailureMessage = "invalid login or password"
)

type ServerHandler struct {
	cfg        *ServerConfig
	authorizer Authorizer
}

func BuildServer(cfg *ServerConfig) *iris.Application {
	//change later
	serv := &ServerHandler{cfg, new(MapAuthorizer)}
	app := iris.New()

	app.Logger().SetLevel(cfg.LogLevel)
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get(fmt.Sprintf("%s/{path:path}", cfg.StaticFilesUrl), serv.HandleStaticFiles)
	app.Post("/signin", serv.AuthHandler)

	app.Get("/", serv.HandleIndex)

	return app
}

func (s *ServerHandler) HandleStaticFiles(ctx iris.Context) {
	cfg := s.cfg
	file := cfg.StaticFileDir + ctx.Params().Get("path")

	if _, err := os.Stat(file); os.IsNotExist(err) {
		ctx.StatusCode(iris.StatusNotFound)
	} else {
		ctx.ServeFile(file, true)
	}
}

func (s *ServerHandler) HandleIndex(ctx iris.Context) {
	ctx.Redirect(s.cfg.StaticFilesUrl + s.cfg.IndexPage)

	//можно и лучше использовать и это, но для этого придется переделать ссылки в верстке
	//лучше в макетах использовать абсолютные ссылки на стили
	//ctx.ServeFile(s.cfg.IndexPage, true)
}

func (s *ServerHandler) AuthHandler(ctx iris.Context) {
	user := new(User)
	response := new(Response)
	ctx.ReadJSON(user)

	if s.authorizer.Authorize(user) {
		response.Status = true
		expiredAt := int32(time.Now().Unix() + DefaultTokenDuration)
		response.Token = NewToken(TmpToken, expiredAt)
	} else {
		response.Status = false
		response.Error = NewError(RequestParsingError, AuthorizationFailureMessage)
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.Write(bytes)
}
