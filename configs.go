package main

import (
	"encoding/json"
	"io/ioutil"
)

//это конфиги для старой сборки сервера, нужно буддет потом структуру поменять
type ServerConfig struct {
	DbConnector        string `json:"dbconnector"`
	DbConnectionString string `json:"connectionstring"`
	DbPort             int    `json:"dbport"`
	DbUser             string `json:"user"`
	DbPassword         string `json:"password"`
	DbName             string `json:"dbname"`
	DbSSL              string `json:"dbssl"`
	CorsAllowedHost    string `json:"corsallowedhost"`
	UseHTTPS           bool   `json:"https"`
	Port               string `json:"port"`
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
