package api

import (
	"log"
	"database/sql"
	"net/http"
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

func ScanUserFromRow(row *sql.Row) (*m.User, error) {
	user := new(m.User)
	err := row.Scan(&user.Id, &user.Nickname, &user.Password, &user.Fullname, &user.Email)

	return user, err
}

//обязательно нужно реализовать
func (u *UserStorage) Add(user *m.User) *ApiResponse {
	result, err := u.db.Exec(
		"INSERT INTO user(nickname, password, email, fullname) VALUES ($1,$2,$3,$4);",
		user.Nickname, user.Password, user.Email, user.Fullname)

	if err != nil {
		log.Println(err)
		return &ApiResponse{
			Code:     409,
			Response: &m.Error{Code: 409, Message: err.Error()}}
	}

	user.Id, _ = result.LastInsertId()
	return &ApiResponse{Code: 201, Response: user}
}

func (u *UserStorage) Remove(user *m.User) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

//untested. Скорее всего не работает
func (u *UserStorage) Update(user *m.User) *ApiResponse {
	row := u.db.QueryRow("SELECT id, nickname, password, fullname, email FROM user WHERE id=$1", user.Id)
	oldUser, err := ScanUserFromRow(row)

	if err != nil {
		log.Println(err)
		return &ApiResponse{
			Code:     http.StatusNotFound,
			Response: &m.Error{Code: http.StatusNotFound, Message: err.Error()},
		}
	}

	if user.Nickname == "" {
		user.Nickname = oldUser.Nickname
	}

	if user.Fullname == "" {
		user.Fullname = oldUser.Fullname
	}

	if user.Password == "" {
		user.Password = oldUser.Password
	}

	if user.Email == "" {
		user.Email = oldUser.Email
	}

	_, err = u.db.Exec("UPDATE user SET nickname=$1, fullname=$2, password=$3, email=$4 WHERE id=$5", user.Nickname, user.Fullname, user.Password, user.Email, user.Id)
	if err != nil {
		log.Println(err)
		return &ApiResponse{
			Code:     http.StatusConflict,
			Response: &m.Error{Code: http.StatusConflict, Message: err.Error()},
		}
	}

	return &ApiResponse{
		Code:     http.StatusOK,
		Response: user,
	}
}

func (u *UserStorage) Get(slug string) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

// UPDATE user SET email = $1, WHERE nickname = $2
// curl -X PUT -d '{"nickname":"asd1, "email":"kek@kek.os"}' localhost:8080/user/
