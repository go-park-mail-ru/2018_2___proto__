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

func NewApiService(connector string, host string,
	port int, user string, password string,
	dbname string, ssl string) (*ApiService, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s",
		host, port, user, dbname, ssl)
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
