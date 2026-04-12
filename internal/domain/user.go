package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	CreatedAt time.Time
}

func (u User) Normalize() User {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	u.Name = strings.TrimSpace(u.Name)
	return u
}

func (u User) Validate() error {
	if u.Email == "" || u.Name == "" {
		return ErrInvalidInput
	}
	return nil
}
