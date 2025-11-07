package model

type AuthCredentials struct {
	Login    string
	Password string
}

type User struct {
	Login    string
	Password string
	UserID   string
}
