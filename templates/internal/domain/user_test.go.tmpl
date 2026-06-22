package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUser_Normalize(t *testing.T) {
	u := User{
		ID:        uuid.MustParse("11111111-1111-1111-1111-111111111111"),
		Email:     "  A@B.COM ",
		Name:      "  Ada  ",
		CreatedAt: time.Unix(0, 0).UTC(),
	}.Normalize()
	if u.Email != "a@b.com" {
		t.Fatalf("email: %q", u.Email)
	}
	if u.Name != "Ada" {
		t.Fatalf("name: %q", u.Name)
	}
}

func TestUser_Validate(t *testing.T) {
	if err := (User{Email: "a@b.com", Name: "n"}).Normalize().Validate(); err != nil {
		t.Fatal(err)
	}
	if err := (User{Email: "", Name: "n"}).Normalize().Validate(); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
	if err := (User{Email: "a@b.com", Name: ""}).Normalize().Validate(); err != ErrInvalidInput {
		t.Fatalf("expected invalid input, got %v", err)
	}
}
