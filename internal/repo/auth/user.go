package auth

import (
	"context"
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
		return nil, err
	}

	row, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	userModel := &repoModel.AuthCredentialsData{}

	err = pgxscan.ScanOne(userModel, row)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(userModel), nil
}
