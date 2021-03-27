package psql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/adrianbrad/ddd-layout/internal/domain"
)

// Ensure the UserService implemented in this package satisfies de UserService interface from the domain.
var _ domain.UserService = (*UserService)(nil)

// UserService represents a service for managing users.
type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// Retrieves a user by ID.
// Returns ENOTFOUND if user does not exist.
// Returns EINTERNAL if operation fails.
func (s *UserService) FindUserByID(ctx context.Context, userID int64) (*domain.User, error) {
   var u domain.User

	err := s.db.
		QueryRowContext(ctx,"SELECT id, name FROM users WHERE id = ?", userID).
		Scan(&u.ID, &u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// return a domain Error containing the ENOTFOUND error code.
			// A more concise error declaration could be domain.Errorf(domain.ENOTFOUND, "user not found")
			return nil, &domain.Error{
				Code:    domain.ENOTFOUND,
				Message: "user not found",
			}
		}

		// if the error is not something that we expect so it's probably a more deep DB error
		// we return the plain error and it will be interpreted as an error with the EINTERNAL error code.
		return nil, err
   }

   return &u, nil
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	// Perform basic field validation
	err := user.Validate()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, "INSERT INTO users(user_id, name) VALUES(?, ?)", user.ID, user.Name)
	if err != nil {
		// As we are not expecting any specific error to be returned.
		// The error received by the caller will always be EINTERNAL
		// and will not be showed to the user.
		return err
	}

	return nil
}