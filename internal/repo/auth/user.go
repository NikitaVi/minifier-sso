package auth

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"github.com/NikitaVi/minifier-sso/internal/repo/converter"
	repoModel "github.com/NikitaVi/minifier-sso/internal/repo/model"
	"github.com/georgysavva/scany/pgxscan"
)

func (r *repository) User(ctx context.Context, login string) (*model.User, error) {
	builder := sq.Select("*").
		PlaceholderFormat(sq.Dollar).
		From(userTableName).Where(sq.Eq{loginColumn: login})

	q, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("Failed to build SQL: %w", err)
	}

	row, err := r.db.Query(ctx, q, args...)
	defer row.Close()
	if err != nil {
		return nil, fmt.Errorf("Failed to execute SQL %s: %w", q, err)
	}

	userModel := &repoModel.AuthCredentialsData{}

	err = pgxscan.ScanOne(userModel, row)
	if err != nil {
		if pgxscan.NotFound(err) {
			return nil, fmt.Errorf("User with login=%s not found: %w", login, err)
		}
		return nil, fmt.Errorf("Failed to scan id user with login=%s: %w", login, err)
	}

	return converter.ToUserFromRepo(userModel), nil
}
