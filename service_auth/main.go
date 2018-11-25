package main

import (
	"log"
	"net"
	"fmt"
	"database/sql"
	"proto-game-server/api"
    m "proto-game-server/models"
	_ "github.com/lib/pq"

    "golang.org/x/net/context"
    "google.golang.org/grpc" 
)

type AuthServer struct {
	Sessions *api.SessionStorage
}


func (as *AuthServer) Auth(ctx context.Context, user *m.User) (*m.SessionId, error) {
	sessionIdString, err := as.Sessions.Create(user)
	sessionId := &m.SessionId{}
	sessionId.Id = sessionIdString

	return sessionId, err
}

func (as *AuthServer) Check(ctx context.Context, sessionId *m.SessionId) (*m.Session, error) {
	return as.Sessions.GetById(sessionId.Id)
}

func (as *AuthServer) LogOut(ctx context.Context, session *m.Session) (*m.Session, error) {
	return as.Sessions.Remove(session)
}

func main() {
	cfg, err := api.LoadConfigs("./data/cfg.json")
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	db, err := sql.Open(cfg.DbConnector, cfg.DbConnectionString)
	if err != nil {
		log.Fatalln(err)
	}

	server := &AuthServer{api.NewSessionStorage(db)}
	serviceServer := grpc.NewServer()
	m.RegisterAuthServer(serviceServer, server)

	fmt.Println(fmt.Sprintf("Starting auth server at %v", cfg.Port))
	serviceServer.Serve(lis)
}