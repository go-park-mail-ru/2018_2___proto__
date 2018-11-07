package models

import (
	"time"
)

type Session struct {
	Id    string `json:"id"`
	Token string `json:"token"`
	//time to live когда сессия умрет
	TTL int64 `json:"ttl"`

	User *User
}

func (s *Session) IsAlive() bool {
	return s.TTL > time.Now().Unix()
}