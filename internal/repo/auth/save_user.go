package auth

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/NikitaVi/minifier-sso/internal/model"
	"github.com/georgysavva/scany/pgxscan"
)

func (r *repository) SaveUser(ctx context.Context, creds *model.AuthCredentials) (string, error) {
	builder := sq.Insert(userTableName).
		PlaceholderFormat(sq.Dollar).
		Columns(loginColumn, passwordColumn).
		Values(creds.Login, creds.Password).
		Suffix("RETURNING id")

	q, args, err := builder.ToSql()
	if err != nil {
		return "", fmt.Errorf("Failed to build SQL: %w", err)
	}

	var id string
	row, err := r.db.Query(ctx, q, args...)
	defer row.Close()
	if err != nil {
		return "", fmt.Errorf("Failed to execute SQL %s: %w", q, err)
	}

	err = pgxscan.ScanOne(&id, row)
	if err != nil {
		return "", fmt.Errorf("Failed to scan id user with login=%s: %w", creds.Login, err)
	}

	return id, nil
}
