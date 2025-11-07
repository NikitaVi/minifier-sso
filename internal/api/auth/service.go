package auth

import (
	"github.com/NikitaVi/minifier-sso/internal/service"
	"github.com/NikitaVi/minifier-sso/pkg/auth_v1"
)

type Implementation struct {
	auth_v1.UnimplementedAuthV1Server
	serv service.AuthService
}

func NewImplementation(serv service.AuthService) *Implementation {
	return &Implementation{
		serv: serv,
	}
}
