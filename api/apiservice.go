package api

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type ApiService struct {
	Users    IUserStorage
	Sessions ISessionStorage
	Scores   IScoreStorage
}

func NewApiService(connector string, connection string) (*ApiService, error) {
	db, err := sql.Open(connector, connection)
	if err != nil {
		return nil, err
	}

	service := &ApiService{
		Users:    NewUserStorage(db),
		Sessions: NewSessionStorage(db),
		Scores:   NewScoreStorage(db),
	}

	return service, nil
}
