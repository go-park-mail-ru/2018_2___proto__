package models

type ScoreRecords struct {
	Count   int           `json:"count"`
	Offset  int           `json:"offset"`
	Records []ScoreRecord `json:"records"`
}
