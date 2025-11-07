package service

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, creds *model.AuthCredentials) (string, error)
	Login(ctx context.Context, creds *model.AuthCredentials) (string, error)
	IsPremium(ctx context.Context, user_id string) (bool, error)
}
