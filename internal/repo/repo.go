package repo

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/model"
)

type AuthRepo interface {
	SaveUser(ctx context.Context, creds *model.AuthCredentials) (string, error)
	User(ctx context.Context, login string) (*model.User, error)
	IsPremium(ctx context.Context, user_id string) (bool, error)
}
