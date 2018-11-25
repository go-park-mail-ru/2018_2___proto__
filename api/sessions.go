package api

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	m "proto-game-server/models"

	"github.com/satori/go.uuid"
)

type ISessionStorage interface {
	//создает сессию для пользователя
	//использовать для авторизации
	Create(user *m.User) (string, error)

	//уничтожает сессиию
	//использовать при выходе из системы
	Remove(session *m.Session) *ApiResponse

	//возвращает сессию и флаг найдена она или нет
	//нужно будет использовать эту функцию при аутентификации
	GetById(id string) (*m.Session, error)
}

type SessionStorage struct {
	db *sql.DB
}

func NewSessionStorage(db *sql.DB) *SessionStorage {
	return &SessionStorage{db: db}
}

// выдача куки при авторизации
func (s *SessionStorage) Create(user *m.User) (string, error) {
	row := s.db.QueryRow(
		"SELECT id, nickname, password, fullname, email, avatar FROM player WHERE nickname=$1",
		user.Nickname)
	expectedUser, err := ScanUserFromRow(row)

	if err != nil {
		log.Print(err)
		return "", err
	}

	if expectedUser.Password != user.Password {
		return "", errors.New("invalid passwred")
	}

	UUID := uuid.NewV4()
	sessionToken := UUID.String()
	var oldToken string
	var ttl int64
	expirationDate := time.Now().Unix() + 86400

	_, err = s.db.Exec("INSERT INTO user_session(token, player_id, expired_date) VALUES ($1, $2, $3);",
		sessionToken, expectedUser.Id, expirationDate)
	if err != nil {
		log.Print(err.Error())
		row = s.db.QueryRow(
			"SELECT token, expired_date FROM user_session WHERE player_id=$1",
			expectedUser.Id)
		err = row.Scan(&oldToken, &ttl)
		if err != nil {
			return "", err
		}
		if ttl < time.Now().Unix() {
			_, err = s.db.Exec("DELETE FROM user_session WHERE token=$1;", oldToken)
			_, err = s.db.Exec("INSERT INTO user_session(token, player_id, expired_date) VALUES ($1, $2, $3);", sessionToken, expectedUser.Id, expirationDate)
			if err != nil {
				return "", errors.New("session is already dead")
			}
		} else {
			sessionToken = oldToken
		}
	}

	log.Println("\nreturned sessionID ", sessionToken)
	return sessionToken, nil
}

func (s *SessionStorage) Remove(session *m.Session) *ApiResponse {
	_, err := s.db.Exec(`DELETE FROM user_session WHERE token=$1`,
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

func (s *SessionStorage) GetById(token string) (*m.Session, error) {
	row := s.db.QueryRow(`SELECT s.id, s.token, s.player_id, s.expired_date, p.id,
	p.nickname, p.password, p.fullname, p.email, p.avatar
		FROM user_session s
		INNER JOIN player p
		on s.player_id = p.id
		WHERE token=$1;`,
		token,
	)

	session, err := ScanSessionFromRow(row)
	return session, err
}
