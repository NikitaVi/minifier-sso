package auth

import (
	"github.com/NikitaVi/minifier-sso/internal/config"
	"github.com/NikitaVi/minifier-sso/internal/repo"
	"github.com/NikitaVi/minifier-sso/internal/service"
)

type serv struct {
	repo repo.AuthRepo
	jwt  config.JWTConfig
}

func NewService(repo repo.AuthRepo, jwtConfig config.JWTConfig) service.AuthService {
	return &serv{repo: repo, jwt: jwtConfig}
}
