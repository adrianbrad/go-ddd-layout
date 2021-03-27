package http

import (
	"encoding/json"
	"github.com/adrianbrad/ddd-layout/internal/domain"
	"net/http"
)

// this is an example of packages defining their our structs for presenting or storing domain type
type user struct {
	UserID int64 `json:"user_id"`//
	Name string `json:"name"`
}


// handleNewUser handles the "POST /user" route.
// It reads & writes data using with HTML or JSON.
func (s *Server) handleNewUser(w http.ResponseWriter, r *http.Request) {
	var u user

	// decode the data into our custom type
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		s.handleErr(w, r, err)
		return
	}

	// pass the data to our user service by converting our custom user type to a domain.User
	err = s.userService.CreateUser(r.Context(), &domain.User{
		ID:  u.UserID ,
		Name: u.Name,
	})
	if err != nil {
		s.handleErr(w, r, err)
		return
	}
}
