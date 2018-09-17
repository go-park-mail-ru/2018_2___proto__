package main

const (
	AuthFailureError    = 1
	RequestParsingError = 2

	DefaultTokenDuration = 8600
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type Token struct {
	Value       string `json:"value"`
	ExpiredDate int32  `json:"expiredat"`
}

type Response struct {
	Status bool   `json:"status"`
	Token  *Token `json:"token,omitempty"`
	Error  *Error `json:"error,omitempty"`
}

func NewToken(value string, expiredDate int32) *Token {
	return &Token{
		Value:       value,
		ExpiredDate: expiredDate,
	}
}

func NewError(code int32, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
