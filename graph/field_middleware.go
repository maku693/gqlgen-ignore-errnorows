package graph

import (
	"context"
	"database/sql"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

func IgnoreErrNoRows(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	res, err := next(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return res, err
}
