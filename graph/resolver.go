//go:generate go run github.com/99designs/gqlgen generate

package graph

import "github.com/maku693/gqlgen-ignore-errnorows/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db db.DB
}

func InitializeResolver(db db.DB) *Resolver {
	return &Resolver{
		db: db,
	}
}
