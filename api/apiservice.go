package api

import (
	"database/sql"
	"fmt"

	"google.golang.org/grpc"

	m "proto-game-server/models"
	_ "github.com/lib/pq"
)

type ApiService struct {
	Sessions m.AuthClient
	Users    IUserStorage
	Scores   IScoreStorage
}

func NewApiService(connector string, connectionString string) (*ApiService, error) {
	fmt.Println(connectionString)

	db, err := sql.Open(connector, connectionString)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	// TODO: fix later
	authClientConn, err := grpc.Dial(
		"127.0.0.1:5050",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	//defer authClientConn.Close()

	service := &ApiService{
		Users:    NewUserStorage(db),
		Sessions: m.NewAuthClient(authClientConn),
		Scores:   NewScoreStorage(db),
	}

	return service, nil
}
