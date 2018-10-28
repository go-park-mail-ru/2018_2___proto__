package api

import (
	"database/sql"
	"net/http"
	"time"
	"log"

	"github.com/satori/go.uuid"
	m "proto-game-server/models"
)

type ISessionStorage interface {
	//создает сессию для пользователя
	//использовать для авторизации
	Create(user *m.User) (string, bool)

	//уничтожает сессиию
	//использовать при выходе из системы
	Remove(session *m.Session) *ApiResponse

	//возвращает сессию и флаг найдена она или нет
	//нужно будет использовать эту функцию при аутентификации
	GetById(id string) (*m.Session, bool)
}

type SessionStorage struct {
	db *sql.DB
}

//нужно реализовать
func NewSessionStorage(db *sql.DB) *SessionStorage {
	return &SessionStorage{db: db}
}

//выдача куки при авторизации
func (s *SessionStorage) Create(user *m.User) (string, bool) {
	row := s.db.QueryRow(
		"SELECT id, nickname, password, fullname, email, avatar FROM player WHERE nickname=$1",
		user.Nickname)
	expectedUser, err := ScanUserFromRow(row)

	if err != nil {
		log.Print(err)
		return "", false
	}

	if expectedUser.Password != user.Password {
		return "", false
	}

	UUID := uuid.NewV4()
	sessionToken := UUID.String()
	expirationDate := time.Now().Unix() + 86400
	
	_, err = s.db.Exec("INSERT INTO user_session(token, player_id, expired_date) VALUES ($1, $2, $3);",
		sessionToken, expectedUser.Id, expirationDate)
	if err != nil {
		print(err.Error())
		row = s.db.QueryRow(
			"SELECT token FROM user_session WHERE player_id=$1",
			expectedUser.Id)
		err = row.Scan(&sessionToken)
		if err != nil {
			return "", false
		}
	}
	return sessionToken, true
}

func (s *SessionStorage) Remove(session *m.Session) *ApiResponse {
	_, err := s.db.Exec(`UPDATE user_session SET ttl=$1 WHERE token=$2`,
		time.Now().Unix(),
		session.Token,
	)

	if err != nil {
		return &ApiResponse{
			Code: http.StatusNotFound,
			Response: &m.Error{
				Code:    http.StatusNotFound,
				Message: err.Error()}}
	}

	return &ApiResponse{
		Code:     http.StatusGone,
		Response: "Session terminated."}
}

func (s *SessionStorage) GetById(token string) (*m.Session, bool) {
	row := s.db.QueryRow(`SELECT user_session.id, user_session.token, user_session.player_id, user_session.expired_date, player.id, player.nickname, player.password, player.fullname, player.email, player.avatar 
	FROM user_session, player 
	WHERE user_session.token=$1;`,
		token,
	)

	session, err := ScanSessionFromRow(row)
	ok := true
	if err != nil {
		print(err.Error())
		ok = false
	}

	return session, ok
}
