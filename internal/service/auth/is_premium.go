package auth

import (
	"context"
)

func (s *serv) IsPremium(ctx context.Context, user_id string) (bool, error) {
	isPremium, err := s.repo.IsPremium(ctx, user_id)
	if err != nil {
		return false, err
	}

	return isPremium, nil
}
