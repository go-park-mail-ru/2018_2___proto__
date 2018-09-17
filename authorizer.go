package main

type Authorizer interface {
	Authorize(*User) bool
}

//make this struc work correctly
type MapAuthorizer struct {
}

func (auth *MapAuthorizer) Authorize(user *User) bool {
	return false
}

func NewMapAuthorizer() *MapAuthorizer {
	return &MapAuthorizer{}
}
