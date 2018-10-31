package api

import (
	"database/sql"
	"errors"
	"net/http"

	m "proto-game-server/models"

	validate "github.com/asaskevich/govalidator"
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

// nice func to remove repeating code
func ThrowAPIError(code int16, message string) *ApiResponse {
	return &ApiResponse{
		Code: int(code),
		Response: &m.Error{
			Code:    code,
			Message: message}}
}

func ValidateUser(user *m.User) (err error) {
	// this defer catches panics from smtp module
	defer func() error {
		if rec := recover(); rec != nil {
			switch x := rec.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown error")
			}
		}
		return err
	}()

	// user fields validation
	_, err = validate.ValidateStruct(user)
	if err != nil {
		return err
	}

	// check if the email is resolvable
	// err = checkmail.ValidateHost(user.Email)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (u *UserStorage) Add(user *m.User) *ApiResponse {
	if err := ValidateUser(user); err != nil {
		return ThrowAPIError(http.StatusBadRequest, err.Error())
	}

	result, err := u.db.Exec(
		"INSERT INTO player(nickname, password, email, fullname) VALUES ($1,$2,$3,$4);",
		user.Nickname, user.Password, user.Email, user.Fullname)

	if err != nil {
		return ThrowAPIError(http.StatusConflict, err.Error())
	}

	user.Id, _ = result.LastInsertId()
	return &ApiResponse{Code: http.StatusCreated, Response: user}
}

// FIXME: remove user's session
func (u *UserStorage) Remove(user *m.User) *ApiResponse {

	// это работает в консоли pgsql, но не работает тут ¯\_(ツ)_/¯
	_, err := u.db.Exec(
		"DELETE FROM player WHERE id=$1;", user.Id)
	print(err.Error())
	if err != nil {
		return ThrowAPIError(http.StatusNotFound, err.Error())
	}

	return &ApiResponse{
		Code:     http.StatusGone,
		Response: "User removed."}
}

// TODO: this funs is untested
func (u *UserStorage) Update(user *m.User) *ApiResponse {
	if _, err := validate.ValidateStruct(user); err != nil {
		return ThrowAPIError(http.StatusBadRequest, err.Error())
	}

	row := u.db.QueryRow("SELECT id, nickname, password, fullname, email, avatar FROM user WHERE id=$1", user.Id)
	oldUser, err := ScanUserFromRow(row)

	if err != nil {
		ThrowAPIError(http.StatusNotFound, err.Error())
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

	if user.Avatar == "" {
		user.Avatar = oldUser.Avatar
	}

	_, err = u.db.Exec("UPDATE user SET nickname=$1, fullname=$2, password=$3, email=$4, avatar=$5 WHERE id=$5",
		user.Nickname, user.Fullname, user.Password, user.Email, user.Id, user.Avatar)
	if err != nil {
		return ThrowAPIError(http.StatusConflict, err.Error())
	}

	return &ApiResponse{
		Code:     http.StatusOK,
		Response: user,
	}
}

// TODO: method for recieving user's info
func (u *UserStorage) Get(slug string) *ApiResponse {
	// TODO: add check for "id" substring in order to search for id

	row := u.db.QueryRow("SELECT id, nickname, email, fullname, avatar FROM player WHERE nickname=$1", slug)
	user := new(m.User)
	err := row.Scan(&user.Id, &user.Nickname, &user.Email, &user.Fullname, &user.Avatar)
	if err != nil {
		return ThrowAPIError(http.StatusNotFound, err.Error())
	}
	return &ApiResponse{
		Code:     http.StatusOK,
		Response: user,
	}
}
