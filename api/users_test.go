package api

import (
	"database/sql"
	m "proto-game-server/models"
	"testing"
)

type RegisterUserCase struct {
	User m.User
	Code int
}

func TestGetUser(t *testing.T) {
	db, _ := sql.Open("postgres", "host=localhost port=5432 user=proto password=proto dbname=proto sslmode=disable")

	us := NewUserStorage(db)

	cases := []RegisterUserCase{
		RegisterUserCase{m.User{Nickname: "Ksenia", Password: "pas@s12www3", Email: "kek@kek.com"}, 201},           // correct
		RegisterUserCase{m.User{Nickname: "Artem", Password: "pswd1@ww23", Email: "kek1@kek.com"}, 201},            // correct
		RegisterUserCase{m.User{Nickname: "Ksenia", Password: "newpawwwss", Email: "kek2@kek.com"}, 409},           // user already exists
		RegisterUserCase{m.User{Nickname: "Ks", Password: "randompass", Email: "kek3@kek.com"}, 409},               // username too short
		RegisterUserCase{m.User{Nickname: "RandomUser", Password: "xz", Email: "kek5@kek.com"}, 400},               // password too short
		RegisterUserCase{m.User{Nickname: "RandomUser2", Password: "*^*GJ*&HChjgj2hg", Email: "kek@kek.com"}, 409}, // email invalid
	}

	for caseNum, item := range cases {
		res := us.Add(&item.User)

		if item.Code != res.Code {
			t.Logf("[%d] unexpected code %d\nreason: %s", caseNum, res.Code, res.Response)
		}
	}

	db.Exec("TRUNCATE TABLE player CASCADE;")

}

/*
	go test -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
*/
