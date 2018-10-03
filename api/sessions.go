package api

import (
	"database/sql"
	"log"
	m "proto-game-server/models"

	"github.com/satori/go.uuid"
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
	db      *sql.DB
	storage map[string]*m.Session
}

//нужно реализовать
func NewSessionStorage(db *sql.DB) *SessionStorage {
	return &SessionStorage{db: db, storage: make(map[string]*m.Session)}
}

//выдача куки при авторизации

func (s *SessionStorage) Create(user *m.User) (string, bool) {
	rows, err := s.db.Query("SELECT password FROM user WHERE nickname=$1", user.Nickname)

	if err != nil {
		log.Print(err)
		return "", false
	}
	defer rows.Close()

	var expectedPassword string
	for rows.Next() {
		err = rows.Scan(&expectedPassword)
		if err = rows.Err(); err != nil {
			log.Print(err)
			return "", false
		}
	}

	if err != nil {
		log.Print(err)
		return "", false
	}

	if expectedPassword != user.Password || expectedPassword == ""{
		return "", false
	}

	UUID, err := uuid.NewV4()
	if err != nil {
		log.Print(err)
		return "", false
	}

	sessionToken := UUID.String()
	s.storage[sessionToken] = &m.Session{User: user}
	return sessionToken, true
}

func (s *SessionStorage) Remove(user *m.Session) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

func (s *SessionStorage) GetById(id string) (*m.Session, bool) {
	return nil, false
}

// curl -d '{"nickname":"asd21","password":"1231",}' localhost:8080/user/signin
