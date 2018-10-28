package game

import (
	"fmt"
)

type Game struct {
	connsChannel chan *Player
}

func NewGame() *Game {
	game := &Game{
		make(chan *Player),
	}

	return game
}

func (g *Game) Start() {
	for conn := range g.connsChannel {
		fmt.Println(conn)
	}
}

func (g *Game) AddPlayer(player *Player) {
	g.connsChannel <- player
}
