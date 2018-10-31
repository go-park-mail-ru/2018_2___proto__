package game

import (
	"proto-game-server/router"
)

//отвечает за мачмэйкинг
type RoomManager struct {
	logger router.ILogger
	queue  chan *Player
}

//достает из очереди пользователей и распределяет их по комнатам
//пока нету обработки преждевременного дисконнекта
func (rm *RoomManager) Start() {
	for {
		player1 := <- rm.queue
		player2 := <- rm.queue

		room := &Room{0, player1, player2, rm}
		go room.Start()
	}
}

func NewRoomManager(logger router.ILogger) *RoomManager {
	roomManager := &RoomManager{
		logger,
		make(chan *Player),
	}

	go roomManager.Start()
	return roomManager
}

func (rm *RoomManager) Queue(player *Player) {
	rm.queue <- player
}
