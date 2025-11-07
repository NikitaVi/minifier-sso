package auth

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/lib/jwt"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *serv) Login(ctx context.Context, creds *model.AuthCredentials) (string, error) {

	userData, err := s.repo.User(ctx, creds.Login)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(creds.Password)); err != nil {
		return "", err
	}

	user := model.User{
		UserID:   userData.UserID,
		Login:    userData.Login,
		Password: userData.Password,
	}

	token, err := jwt.NewToken(user, s.jwt)
	if err != nil {
		return "", err
	}

	return token, nil
}
