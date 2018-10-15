package models

type Error struct {
	Code    int16  `json:"code"`
	Message string `json:"msg"`
}
