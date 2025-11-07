package auth

import (
	"github.com/NikitaVi/minifier-sso/internal/repo"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) repo.AuthRepo {
	return &repository{db: db}
}
