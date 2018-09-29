package models

type Session struct {
	Id string `json:"id"`

	//time to live сколько сессия еще будет жить
	TTL int32 `json:"ttl"`

	User *User
}