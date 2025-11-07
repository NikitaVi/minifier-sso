package auth

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *serv) Register(ctx context.Context, creds *model.AuthCredentials) (string, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	userCreds := &model.AuthCredentials{Login: creds.Login, Password: string(passHash)}

	id, err := s.repo.SaveUser(ctx, userCreds)
	if err != nil {
		return "", err
	}
	return id, nil
}
