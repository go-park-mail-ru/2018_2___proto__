package game

import (
	ws "github.com/gorilla/websocket"
	m "proto-game-server/models"
)

type Player struct {
	session *m.Session
	conn    *ws.Conn
}

func NewPlayer(session *m.Session, conn *ws.Conn) *Player {
	return &Player{session, conn}
}
