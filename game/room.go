package game

/*
	I DID NOT WRITE THIS
	I DID NOT
	IT'S BULLSHIT
	*falling keyboard noises*
	I DID NOOOOOOOT
	OH, HI DEADLINE
*/

//собственно здесь будет вся игра проходить
//получение комманд от пользователей
//слежение за временем и тд
type Room struct {
	id          int
	player1     *Player
	player2     *Player
	roomManager *RoomManager
}

func (r *Room) ListenPlayer(player *Player) {
	player.conn.WriteJSON(`{"command":123;"value":"asd"}`)

	for {
		command, err := player.ReadCommand()

		if err != nil {
			r.roomManager.logger.Error(err)
			continue
		}

		switch command.Type {
		case CMD_FINDROOM:
			break

		case CMD_DISCONNECT:
			break

		default:
			break
		}

		r.roomManager.logger.Notice(command)
	}
}

//здесь происходит обработка комманд пользователей
//позже нужно будет добавить спец таймер для определения хода
func (r *Room) Start() {

}