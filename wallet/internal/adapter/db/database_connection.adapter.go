package database

import (
	"context"
	"wallet/internal/interfaces"

	"github.com/jackc/pgx/v5"
)

type DatabaseConnectionAdapter struct {
	Conn *pgx.Conn
}

type pgxRowsAdapter struct {
	pgx.Rows
}

func (r *pgxRowsAdapter) Next() bool                     { return r.Rows.Next() }
func (r *pgxRowsAdapter) Scan(dest ...interface{}) error { return r.Rows.Scan(dest...) }
func (r *pgxRowsAdapter) Close()                         { r.Rows.Close() }
func (r *pgxRowsAdapter) Err() error                     { return r.Rows.Err() }

func (p *DatabaseConnectionAdapter) Query(ctx context.Context, sql string, args ...interface{}) (interfaces.Rows, error) {
	rows, err := p.Conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &pgxRowsAdapter{rows}, nil
}

func (p *DatabaseConnectionAdapter) Exec(ctx context.Context, sql string, args ...interface{}) (interfaces.Result, error) {
	return p.Conn.Exec(ctx, sql, args...)
}
