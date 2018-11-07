package api

import (
	"database/sql"
	"net/http"
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

func ReadScoreRecords(rows *sql.Rows, usedOffset int) (*m.ScoreRecords, error) {
	scores := make([]*m.ScoreRecord, 0)

	for rows.Next() {
		scoreRecord, err := ScanScoreFromRow(rows)
		if err != nil {
			return nil, err
		}

		scores = append(scores, scoreRecord)
	}

	scoreRecords := &m.ScoreRecords{
		Count:   len(scores),
		Offset:  usedOffset,
		Records: scores,
	}

	return scoreRecords, nil
}

func (s *ScoreStorage) Get(offset int, limit int) *ApiResponse {
	rows, err := s.db.Query(`WITH scores AS (
		SELECT id, score, player_id 
		FROM score
		ORDER BY score DESC, id ASC
		LIMIT $1 OFFSET $2
	)
	
	SELECT s.id, s.score, p.nickname
	FROM scores AS s
	INNER JOIN player AS p ON p.id=s.player_id`,
		limit,
		offset,
	)

	if err != nil {
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	scoresRecords, err := ReadScoreRecords(rows, offset)
	if err != nil {
		return &ApiResponse{Code: http.StatusBadRequest, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: scoresRecords}
}
