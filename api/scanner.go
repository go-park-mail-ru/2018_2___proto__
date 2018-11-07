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

func ScanSessionFromRow(row IRow) (*m.Session, error) {
	session := new(m.Session)
	user := new(m.User)
	session.User = user
	err := row.Scan(&session.Id, &session.Token, &session.User.Id, &session.TTL,
		&session.User.Id, &session.User.Nickname, &session.User.Password,
		&session.User.Fullname, &session.User.Email, &session.User.Avatar)
	return session, err
}

func ScanScoreFromRow(row IRow) (*m.ScoreRecord, error) {
	scoreRecord := new(m.ScoreRecord)
	err := row.Scan(&scoreRecord.Id, &scoreRecord.Score, &scoreRecord.Nickname)
	return scoreRecord, err
}