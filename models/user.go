package models

type User struct {
	Id       int32  `json:"id,omitempty"`
	Nickname string `json:"nickname"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}