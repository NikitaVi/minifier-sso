package auth

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"github.com/NikitaVi/minifier-sso/pkg/auth_v1"
)

func (i *Implementation) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	userCreds := &model.AuthCredentials{Login: req.Login, Password: req.Password}
	token, err := i.serv.Login(ctx, userCreds)
	if err != nil {
		return nil, err
	}

	return &auth_v1.LoginResponse{Token: token}, nil
}
