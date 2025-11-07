package auth

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"github.com/NikitaVi/minifier-sso/pkg/auth_v1"
)

func (i *Implementation) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	userCreds := &model.AuthCredentials{Login: req.Login, Password: req.Password}

	userId, err := i.serv.Register(ctx, userCreds)
	if err != nil {
		return nil, err
	}

	return &auth_v1.RegisterResponse{UserGuid: userId}, nil
}
