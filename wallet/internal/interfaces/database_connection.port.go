package interfaces

import (
	"context"
)

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close()
	Err() error
}

type Result interface{}

type DatabaseConnection interface {
	Query(ctx context.Context, sql string, args ...interface{}) (Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (Result, error)
}
