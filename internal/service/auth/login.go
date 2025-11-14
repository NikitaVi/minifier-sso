package auth

import (
	"context"
	"fmt"
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
		return "", fmt.Errorf("Failed to compare password: %w", err)
	}

	user := model.User{
		UserID:   userData.UserID,
		Login:    userData.Login,
		Password: userData.Password,
	}

	token, err := jwt.NewToken(user, s.jwt)
	if err != nil {
		return "", fmt.Errorf("Failed to create token: %w", err)
	}

	return token, nil
}
