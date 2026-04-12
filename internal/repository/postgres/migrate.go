package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Migrate(ctx context.Context, pool *pgxpool.Pool) error {
	const ddl = `
create table if not exists users (
	id uuid primary key,
	email text not null unique,
	name text not null,
	created_at timestamptz not null default now()
);
`
	if _, err := pool.Exec(ctx, ddl); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}
