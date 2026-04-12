package user

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/kikplate-plates/go-clean-architecture-starter/internal/domain"
)

type fakeGetRepo struct {
	user domain.User
	err  error
}

func (f *fakeGetRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {
	return domain.User{}, errors.New("not_used")
}

func (f *fakeGetRepo) GetByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	if f.err != nil {
		return domain.User{}, f.err
	}
	return f.user, nil
}

func TestGetUser_Execute_InvalidID(t *testing.T) {
	uc := GetUser{Repo: &fakeGetRepo{}}
	_, err := uc.Execute(context.Background(), uuid.Nil)
	if !errors.Is(err, domain.ErrInvalidInput) {
		t.Fatalf("expected invalid input, got %v", err)
	}
}

func TestGetUser_Execute_NotFound(t *testing.T) {
	uc := GetUser{Repo: &fakeGetRepo{err: domain.ErrNotFound}}
	_, err := uc.Execute(context.Background(), uuid.MustParse("22222222-2222-2222-2222-222222222222"))
	if !errors.Is(err, domain.ErrNotFound) {
		t.Fatalf("expected not found, got %v", err)
	}
}

func TestGetUser_Execute_OK(t *testing.T) {
	id := uuid.MustParse("33333333-3333-3333-3333-333333333333")
	want := domain.User{ID: id, Email: "x@y.z", Name: "x"}
	uc := GetUser{Repo: &fakeGetRepo{user: want}}
	out, err := uc.Execute(context.Background(), id)
	if err != nil {
		t.Fatal(err)
	}
	if out.ID != want.ID || out.Email != want.Email {
		t.Fatalf("unexpected user: %+v", out)
	}
}
