package domain

import (
	"context"
)

// User represents a user in the system.
type User struct {
	ID int64
	Name string
}

// Validate returns an error if the user contains invalid fields.
// This only performs basic validation.
func (u *User) Validate() error {
	if u.Name == "" {
		return Errorf(EINVALID, "user name is required.")
	}
	return nil
}

// UserService represents a service for managing users.
type UserService interface {
	// Creates a new user
	// Returns EINTERNAL if operation fails.
	CreateUser(ctx context.Context, user *User) error

	// Retrieves a user by ID.
	// Returns ENOTFOUND if user does not exist.
	// Returns EINTERNAL if operation fails.
	FindUserByID(ctx context.Context, userID int64) (*User, error)
}