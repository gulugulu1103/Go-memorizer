package model

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

// UserService defines the interface for user service
type UserService interface {
	Get(ctx context.Context, uid uuid.UUID) (*User, error) // Get a user by uid
}

// UserRepository defines the interface for user repository
type UserRepository interface {
	FindByID(uid uuid.UUID) (*User, error) // Get a user by uid
}
