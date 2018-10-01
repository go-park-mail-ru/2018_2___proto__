package models

type User struct {
	Id       int64  `json:"id,omitempty"`
	Nickname string `json:"nickname"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}