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
	row, err := s.db.Query("SELECT password FROM user WHERE nickname=$1", user.Nickname)

	if err != nil {
		log.Fatal(err)
		return "", false
	}
	defer row.Close()

	var expectedPassword string
	for row.Next() {
		err = row.Scan(&expectedPassword)
		if err = row.Err(); err != nil {
			log.Fatal(err)
			return "", false
		}
	}

	if err != nil {
		log.Fatal(err)
		return "", false
	}

	if expectedPassword != user.Password {
		return "", false
	}

	UUID, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
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
