package api

import (
	m "proto-game-server/models"
)

type ISessionStorage interface {
	//создает сессию для пользователя
	//использовать для авторизации
	Create(user *m.User) string

	//уничтожает сессиию
	//использовать при выходе из системы
	Remove(session *m.Session) *ApiResponse

	//возвращает сессию и флаг найдена она или нет
	//нужно будет использовать эту функцию при аутентификации
	GetById(id string) (*m.Session, bool)
}

type SessionStorage struct {
}

//нужно реализовать
func NewSessionStorage() *SessionStorage {
	//тут должна проходить инициализация хранилища сессий
	return &SessionStorage{}
}

//выдача куки при авторизации
//нужно реализовать
func (s *SessionStorage) Create(user *m.User) string {
	return "sessionId"
}

func (s *SessionStorage) Remove(user *m.Session) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

func (s *SessionStorage) GetById(id string) (*m.Session, bool) {
	return nil, false
}
