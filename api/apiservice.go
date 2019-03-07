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

func NewApiService(cfg *ServerConfig) (*ApiService, error) {
	fmt.Println(cfg.DbConnectionString)

	db, err := sql.Open(cfg.DbConnector, cfg.DbConnectionString)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	// TODO: fix later
	authClientConn, err := grpc.Dial(
		cfg.AuthServiceHost,
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