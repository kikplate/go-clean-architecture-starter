package postgres

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kikplate-plates/go-clean-architecture-starter/internal/domain"
)

func openPool(t *testing.T) *UserRepository {
	t.Helper()
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set")
	}
	ctx := context.Background()
	pool, err := Connect(ctx, dsn)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	if err := Migrate(ctx, pool); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), "truncate table users")
	})
	return &UserRepository{Pool: pool}
}

func TestUserRepository_CreateAndGet(t *testing.T) {
	repo := openPool(t)
	ctx := context.Background()
	u := domain.User{
		ID:        uuid.New(),
		Email:     "repo+" + uuid.New().String() + "@example.com",
		Name:      "Repo",
		CreatedAt: time.Now().UTC().Truncate(time.Millisecond),
	}.Normalize()

	created, err := repo.Create(ctx, u)
	if err != nil {
		t.Fatal(err)
	}
	if created.ID != u.ID {
		t.Fatalf("id mismatch")
	}

	got, err := repo.GetByID(ctx, u.ID)
	if err != nil {
		t.Fatal(err)
	}
	if got.Email != u.Email || got.Name != u.Name {
		t.Fatalf("unexpected: %+v", got)
	}
}

func TestUserRepository_GetByID_NotFound(t *testing.T) {
	repo := openPool(t)
	_, err := repo.GetByID(context.Background(), uuid.New())
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestUserRepository_Create_Conflict(t *testing.T) {
	repo := openPool(t)
	ctx := context.Background()
	email := "dup+" + uuid.New().String() + "@example.com"
	base := domain.User{
		ID:        uuid.New(),
		Email:     email,
		Name:      "a",
		CreatedAt: time.Now().UTC().Truncate(time.Millisecond),
	}.Normalize()
	if _, err := repo.Create(ctx, base); err != nil {
		t.Fatal(err)
	}
	dup := domain.User{
		ID:        uuid.New(),
		Email:     email,
		Name:      "b",
		CreatedAt: time.Now().UTC().Truncate(time.Millisecond),
	}.Normalize()
	_, err := repo.Create(ctx, dup)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, domain.ErrConflict) {
		t.Fatalf("expected conflict, got %v", err)
	}
}
