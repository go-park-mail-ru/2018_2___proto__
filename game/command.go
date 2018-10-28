package game

const (
	CMD_FINDROOM   = 1
	CMD_DISCONNECT = 2
)

//комманды, посылаемые пользователем в игровую комнату
//формат будет менятся в процессе разработки
type Command struct {
	Type  int    `json:"type"`
	Value string `json:"value"`
}
