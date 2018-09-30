package api

import (
	"database/sql"
	m "proto-game-server/models"
)

type IUserStorage interface {
	Add(user *m.User) *ApiResponse

	Remove(user *m.User) *ApiResponse

	Update(user *m.User) *ApiResponse

	Get(slug string) *ApiResponse
}

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db}
}

//обязательно нужно реализовать
func (u *UserStorage) Add(user *m.User) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

func (u *UserStorage) Remove(user *m.User) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

func (u *UserStorage) Update(user *m.User) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

func (u *UserStorage) Get(slug string) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}