package main

import (
	"encoding/json"
	"io/ioutil"
)

//это конфиги для старой сборки сервера, нужно буддет потом структуру поменять
type ServerConfig struct {
	IndexPage      string `json:"indexpage"`
	StaticFilesUrl string `json:"staticfilesurl"`
	StaticFileDir  string `json:"staticfilesdir"`
	ServerPort     string `json:"serverport"`
	LogLevel       string `json:"loglevel"`
}

func LoadConfigs(path string) (*ServerConfig, error) {
	cfg := new(ServerConfig)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
