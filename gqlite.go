// Package gqlite provides a GoHub database implementation backed by GraphQLite.
package gqlite

import (
	"context"
	"github.com/mughub/mughub/db"
	"github.com/spf13/viper"
	"io"
)

// DB is a database implementation backed by GraphQLite
type DB struct {
	// TODO: Add GraphQLite database here
}

// Init initializes a GraphQLite database.
func (d *DB) Init(schema io.Reader, cfg *viper.Viper) (err error) {
	return
}

// Do executes a GraphQL request against the GraphQLite database.
func (d *DB) Do(ctx context.Context, req string, vars map[string]interface{}) (r *db.Result) {
	return
}
