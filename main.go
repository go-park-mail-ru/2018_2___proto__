package main

import (
	"github.com/kataras/iris"
)

func main() {
	cfg, err := LoadConfigs("./cfg.json")
	if err != nil {
		panic("cfg file not found")
	}

	app := BuildServer(cfg)
	app.Run(iris.Addr(cfg.ServerPort))
}
