package api

import (
	"database/sql"
	"fmt"
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
	fmt.Println(user.Nickname, user.Password, user.Email, user.Fullname)
	_, err := u.db.Query("INSERT INTO user(nickname, password, email, fullname) VALUES (\"$1\",\"$2\",\"$3\",\"$4\");",
		user.Nickname, user.Password, user.Email, user.Fullname)
	raws, _ := u.db.Query("SELECT * FROM proto.sqlite_master WHERE type='table'")

	fmt.Println(err)
	fmt.Println(raws)
	if err != nil {
		return &ApiResponse{
			Code:     409,
			Response: &m.Error{Code: 409, Message: "Username is already occupied"}}
	}
	return &ApiResponse{Code: 201}
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
