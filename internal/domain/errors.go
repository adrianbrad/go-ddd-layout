package domain

// Comments found in this file are mostly taken from here: https://github.com/benbjohnson/wtf/blob/main/error.go

import (
	"errors"
	"fmt"
)

// Application error codes.
// These are not actual errors, they are just an property to be used along with a domain Error.
//
// The format used (all caps) is taken from the linux error codes format:
// /usr/include/asm/errno.h
// #define EPERM            1      /* Operation not permitted */
// #define ENOENT           2      /* No such file or directory */
// #define ESRCH            3      /* No such process */
// #define EINTR            4      /* Interrupted system call */
//
// Note: These are meant to be generic an map closely to the HTTP error codes.
//
const (
	EINTERNAL = "internal"
	ENOTFOUND = "not_found"
	EINVALID = "invalid"
)
/*
Error represents an application-specific error. Application errors can be
unwrapped by the caller to extract out the code & message.

Any non-application error (such as a unexpected sql error) should be reported as an
EINTERNAL error and the human user should only see "Internal error" as the
message. These low-level internal error details should only be logged and
reported to the operator of the application (not the end user).

Example:

package psql

import (
	"database/sql"
	"errors"
	"github.com/adrianbrad/ddd-layout/internal/domain"
)

func (s *UserService) FindUser(userID int64) (*domain.User, error) {
	var u domain.User

	err := s.db.
		QueryRow("SELECT id, name FROM users WHERE id = ?", userID).
		Scan(&u.ID, &u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// if this is an expected error with decorate it with a domain error code.
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
*/
type Error struct {
	// Machine-readable error code.
	Code string

	// Human-readable error message.
	Message string
}

// Error implements the error interface. Not used by the application otherwise.
func (e *Error) Error() string {
	return fmt.Sprintf("wtf error: code=%s message=%s", e.Code, e.Message)
}

// ErrorCode unwraps an application error and returns its code.
// Non-application errors always return EINTERNAL.
func ErrorCode(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Code
	}
	return EINTERNAL
}

// ErrorMessage unwraps an application error and returns its message.
// Non-application errors always return "Internal error".
func ErrorMessage(err error) string {
	var e *Error
	if err == nil {
		return ""
	} else if errors.As(err, &e) {
		return e.Message
	}
	return "Internal error."
}

// Errorf is a helper function to return an Error with a given code and formatted message.
func Errorf(code string, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}