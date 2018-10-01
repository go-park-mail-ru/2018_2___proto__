package api

import (
	"database/sql"
	m "proto-game-server/models"
)

type IScoreStorage interface {
	Get(offset int, limit int) *ApiResponse
}

type ScoreStorage struct {
	db *sql.DB
}

func NewScoreStorage(db *sql.DB) *ScoreStorage {
	return &ScoreStorage{db}
}

func (s *ScoreStorage) Get(offset int, limit int) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}