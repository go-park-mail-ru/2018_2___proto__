package game

import (
	"proto-game-server/router"
)

type Game struct {
	playerChannel chan *Player
	roomManager   *RoomManager
	logger        router.ILogger
}

func NewGame(logger router.ILogger) *Game {
	game := &Game{
		make(chan *Player),
		NewRoomManager(logger),
		logger,
	}

	return game
}

func (g *Game) Start() {
	for player := range g.playerChannel {
		g.roomManager.Queue(player)
	}
}

func (g *Game) AddPlayer(player *Player) {
	g.playerChannel <- player
}
