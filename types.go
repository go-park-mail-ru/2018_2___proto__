package main

type UserFull struct {
	Id       int32  `json:"id,omitempty"`
	Nickname string `json:"nickname"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Fullname string `json:"fullname,omitempty"`
}

type ScoreRecord struct {
	Id       int32
	Score    int32  `json:"score"`
	Nickname string `json:"nickname"`
}

type Session struct {
	Id string `json:"id"`

	//time to live сколько сессия еще будет жить
	TTL int32 `json:"ttl"`
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"msg"`
}
