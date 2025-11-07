package auth

import (
	"context"
	"github.com/NikitaVi/minifier-sso/pkg/auth_v1"
)

func (i *Implementation) IsPremium(ctx context.Context, req *auth_v1.IsPremiumRequest) (*auth_v1.IsPremiumResponse, error) {
	isPremium, err := i.serv.IsPremium(ctx, req.UserGuid)
	if err != nil {
		return nil, err
	}

	return &auth_v1.IsPremiumResponse{IsPremium: isPremium}, nil
}
