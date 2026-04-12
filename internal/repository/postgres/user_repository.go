package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kikplate-plates/go-clean-architecture-starter/internal/domain"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func (r UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	const q = `
insert into users (id, email, name, created_at)
values ($1, $2, $3, $4)
returning id, email, name, created_at
`
	row := r.Pool.QueryRow(ctx, q, user.ID, user.Email, user.Name, user.CreatedAt)
	var out domain.User
	if err := row.Scan(&out.ID, &out.Email, &out.Name, &out.CreatedAt); err != nil {
		if isUniqueViolation(err) {
			return domain.User{}, domain.ErrConflict
		}
		return domain.User{}, fmt.Errorf("insert user: %w", err)
	}
	return out, nil
}

func (r UserRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	const q = `select id, email, name, created_at from users where id = $1`
	row := r.Pool.QueryRow(ctx, q, id)
	var out domain.User
	if err := row.Scan(&out.ID, &out.Email, &out.Name, &out.CreatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, domain.ErrNotFound
		}
		return domain.User{}, fmt.Errorf("select user: %w", err)
	}
	return out, nil
}

func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}
