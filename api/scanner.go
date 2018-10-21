package api

import (
	m "proto-game-server/models"
)

type IRow interface {
	Scan(dest ...interface{}) error
}


func ScanUserFromRow(row IRow) (*m.User, error) {
	user := new(m.User)
	err := row.Scan(&user.Id, &user.Nickname, &user.Password, &user.Fullname, &user.Email, &user.Avatar)

	return user, err
}