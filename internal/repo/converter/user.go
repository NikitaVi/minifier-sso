package converter

import (
	"github.com/NikitaVi/minifier-sso/internal/model"
	repoModel "github.com/NikitaVi/minifier-sso/internal/repo/model"
)

func ToUserFromRepo(user *repoModel.AuthCredentialsData) *model.User {
	return &model.User{
		UserID:   user.ID,
		Login:    user.Login,
		Password: user.Password,
	}
}
