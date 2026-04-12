package user

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kikplate-plates/go-clean-architecture-starter/internal/domain"
)

type fakeUserRepo struct {
	created domain.User
	err     error
}

func (f *fakeUserRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	if f.err != nil {
		return domain.User{}, f.err
	}
	return f.created, nil
}

func (f *fakeUserRepo) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return domain.User{}, errors.New("not_used")
}

func TestCreateUser_Execute_InvalidInput(t *testing.T) {
	uc := CreateUser{Repo: &fakeUserRepo{}}
	_, err := uc.Execute(context.Background(), CreateUserInput{Email: "", Name: "n"})
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected invalid input, got %v", err)
	}
}

func TestCreateUser_Execute_OK(t *testing.T) {
	id := uuid.New()
	at := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	repo := &fakeUserRepo{
		created: domain.User{ID: id, Email: "a@b.com", Name: "ab", CreatedAt: at},
	}
	uc := CreateUser{Repo: repo}
	out, err := uc.Execute(context.Background(), CreateUserInput{Email: "A@B.com", Name: " ab "})
	if err != nil {
		t.Fatal(err)
	}
	if out.Email != "a@b.com" || out.Name != "ab" {
		t.Fatalf("unexpected normalization: %+v", out)
	}
}

func TestCreateUser_Execute_RepoError(t *testing.T) {
	uc := CreateUser{Repo: &fakeUserRepo{err: domain.ErrConflict}}
	_, err := uc.Execute(context.Background(), CreateUserInput{Email: "a@b.com", Name: "n"})
	if !errors.Is(err, domain.ErrConflict) {
		t.Fatalf("expected conflict, got %v", err)
	}
}
