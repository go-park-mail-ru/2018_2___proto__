package main

type Authorizer interface {
	Authorize(*User) bool
}

type MapAuthorizer struct {

}

func (auth *MapAuthorizer) Authorize(user *User) bool {
	return false
} 