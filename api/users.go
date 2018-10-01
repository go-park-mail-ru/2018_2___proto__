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
	_, err := u.db.Exec(
		"INSERT INTO user(nickname, password, email, fullname) VALUES ($1,$2,$3,$4);",
		user.Nickname, user.Password, user.Email, user.Fullname)
	fmt.Println(user.Nickname, user.Password, user.Email, user.Fullname)
	if err != nil {
		fmt.Println(err.Error())
		return &ApiResponse{
			Code:     409,
			Response: &m.Error{Code: 409, Message: err.Error()}}
	}
	return &ApiResponse{Code: 201}
}

func (u *UserStorage) Remove(user *m.User) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

//untested. Скорее всего не работает
func (u *UserStorage) Update(user *m.User) *ApiResponse {
	// return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
	if user.Nickname == "" {
		return &ApiResponse{
			Code:     400,
			Response: &m.Error{Code: 400, Message: "Omitted username"}}
	}
	fmt.Println("user.Email")
	if user.Email != "" {
		_, err := u.db.Exec(
			"UPDATE user SET email = $1 WHERE nickname = $2",
			user.Email, user.Nickname)
		if err != nil {
			return &ApiResponse{
				Code:     400,
				Response: &m.Error{Code: 400, Message: err.Error()}}
		}
	}
	if user.Password != "" {
		_, err := u.db.Exec(
			"UPDATE user SET password = $1 WHERE nickname = $2",
			user.Password, user.Nickname)
		if err != nil {
			return &ApiResponse{
				Code:     400,
				Response: &m.Error{Code: 400, Message: err.Error()}}
		}
	}
	if user.Fullname != "" {
		_, err := u.db.Exec(
			"UPDATE user SET fullname = $1 WHERE nickname = $2",
			user.Fullname, user.Nickname)
		if err != nil {
			return &ApiResponse{
				Code:     400,
				Response: &m.Error{Code: 400, Message: err.Error()}}
		}
	}

	return &ApiResponse{Code: 200}
}

/* func (u *UserStorage) Update(user *m.User) *ApiResponse {
	// return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
	if user.Nickname == "" {
		return &ApiResponse{
			Code:     400,
			Response: &m.Error{Code: 400, Message: "Omitted username"}}
	}

	row, err := u.db.Query("SELECT email, fullname, password FROM user WHERE nickname=$1", user.Nickname)

	if err != nil {
		log.Fatal(err)
		return &ApiResponse{
			Code:     409,
			Response: &m.Error{Code: 409, Message: err.Error()}}
	}
	defer row.Close()

	var oldUser m.User
	for row.Next() {
		err = row.Scan(&oldUser.Email, &oldUser.Fullname, &oldUser.Password)
		if err != nil {
			log.Fatal(err)
			return &ApiResponse{
				Code:     409,
				Response: &m.Error{Code: 409, Message: err.Error()}}
		}
	}
	if err = row.Err(); err != nil {
		log.Fatal(err)
		return &ApiResponse{
			Code:     409,
			Response: &m.Error{Code: 409, Message: err.Error()}}
	}

	// query := "UPDATE user SET "
	// counter := 1

	fmt.Println(oldUser)
	if user.Email != "" {
		_, err := u.db.Exec(
			"UPDATE user SET email = $1 WHERE nickname = $2",
			user.Email, user.Nickname)
		if err != nil {
			return &ApiResponse{
				Code:     400,
				Response: &m.Error{Code: 400, Message: err.Error()}}
		}
	}

	return &ApiResponse{Code: 200}
} */

func (u *UserStorage) Get(slug string) *ApiResponse {
	return &ApiResponse{Code: 400, Response: &m.Error{1, "unimplemented api"}}
}

// UPDATE user SET email = $1, WHERE nickname = $2
// curl -X PUT -d '{"nickname":"asd1, "email":"kek@kek.os"}' localhost:8080/user/
