package models

type Session struct {
	Id    string `json:"id"`
	Token string `json:"token"`
	//time to live когда сессия умрет
	TTL int64 `json:"ttl"`

	User *User
}
