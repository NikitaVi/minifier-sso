package auth

import (
	"context"
	"github.com/NikitaVi/minifier-sso/internal/logger"
	"github.com/NikitaVi/minifier-sso/pkg/auth_v1"
)

func (i *Implementation) IsPremium(ctx context.Context, req *auth_v1.IsPremiumRequest) (*auth_v1.IsPremiumResponse, error) {
	isPremium, err := i.serv.IsPremium(ctx, req.UserGuid)
	if err != nil {
		logger.Error("IsPremium request failed (id=%s): %v ", req.UserGuid, err)
		return nil, err
	}

	return &auth_v1.IsPremiumResponse{IsPremium: isPremium}, nil
}
