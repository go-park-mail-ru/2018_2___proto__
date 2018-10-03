package api

import (
	"database/sql"
	m "proto-game-server/models"
)

//хардкод лидеров
var leaders = []*m.ScoreRecord {
	&m.ScoreRecord {
		Id : 1,
		Score : 1000,
		Nickname : "user1",
	},
	&m.ScoreRecord {
		Id : 2,
		Score : 5000,
		Nickname : "user2",
	},
	&m.ScoreRecord {
		Id : 5,
		Score : 100,
		Nickname : "user5",
	},
	&m.ScoreRecord {
		Id : 3,
		Score : 10,
		Nickname : "user3",
	},
	&m.ScoreRecord {
		Id : 4,
		Score : 0,
		Nickname : "user4",
	},
}

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
	l := len(leaders)
	limit = offset + l

	if limit > l {
		limit = l
	}

	if offset < 0 || offset > l {
		offset = 0
	}

	records := &m.ScoreRecords {
		Count : limit,
		Offset : offset,
		Records : leaders[offset:offset+limit],
	}

	return &ApiResponse{Code: 200, Response: records}
}