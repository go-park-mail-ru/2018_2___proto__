package models

type ScoreRecord struct {
	Id       int32	`json:"id"`
	Score    int32  `json:"score"`
	Nickname string `json:"nickname"`
}