package api

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type ApiService struct {
	Users    IUserStorage
	Sessions ISessionStorage
	Scores   IScoreStorage
}

func NewApiService(connector string, connectionString string) (*ApiService, error) {
	fmt.Println(connectionString)

	db, err := sql.Open(connector, connectionString)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	service := &ApiService{
		Users:    NewUserStorage(db),
		Sessions: NewSessionStorage(db),
		Scores:   NewScoreStorage(db),
	}

	return service, nil
}
