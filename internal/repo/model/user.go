package model

import "time"

type IsPremiumData struct {
	EndDate time.Time `db:"end_date"`
	Active  bool      `db:"is_active"`
}

type AuthCredentialsData struct {
	ID       string `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
