package auth

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/NikitaVi/minifier-sso/internal/repo/model"
	"github.com/georgysavva/scany/pgxscan"
	"time"
)

func (r *repository) IsPremium(ctx context.Context, userId string) (bool, error) {
	builder := sq.Select(activeColumn, endDateColumn).
		PlaceholderFormat(sq.Dollar).
		From(userPremiumTableName).Where(sq.Eq{userIdColumn: userId})

	q, args, err := builder.ToSql()
	if err != nil {
		return false, err
	}

	var data model.IsPremiumData
	row, err := r.db.Query(ctx, q, args...)
	if err != nil {
		return false, err
	}

	err = pgxscan.ScanOne(&data, row)
	if err != nil {
		if pgxscan.NotFound(err) {
			return false, nil
		}
		return false, err
	}

	now := time.Now().UTC()

	isActive := data.Active && data.EndDate.After(now)
	return isActive, nil
}
