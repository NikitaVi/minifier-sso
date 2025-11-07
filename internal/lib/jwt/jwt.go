package jwt

import (
	"github.com/NikitaVi/minifier-sso/internal/config"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewToken(user model.User, jwtConfig config.JWTConfig) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.UserID,
		"login": user.Login,
		"exp":   time.Now().Add(jwtConfig.TTL()).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtConfig.Secret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
