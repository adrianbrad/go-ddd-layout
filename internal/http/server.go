package http

import (
	"encoding/json"
	"github.com/adrianbrad/ddd-layout/internal/domain"
	"net/http"
)

type Server struct {
	pubSub domain.PubSub
	userService domain.UserService
}

func NewServer(pubSub domain.PubSub, userService domain.UserService) *Server {
	return &Server{
		pubSub:      pubSub,
		userService: userService,
	}
}

type errorResponse struct {
	Code string `code:"code"`
	Message string `json:"message"`
}

func (s *Server) handleErr(w http.ResponseWriter, r *http.Request, err error) {
	// extract the error code
	code := domain.ErrorCode(err)

	// if we have an internal error we log, report and hide the details from the user
	if code == domain.EINTERNAL {
		s.logError(err)
		s.reportErr(err)
		err = json.NewEncoder(w).Encode(errorResponse{
			Code:    code,
			Message: "internal server error.",
		})
		if err != nil {
			s.logError(err)
		}
		return
	}

	// if the error is not internal present the user with the error message
	message := domain.ErrorMessage(err)
	err = json.NewEncoder(w).Encode(errorResponse{
		Code:    code,
		Message: message,
	})
	if err != nil {
		s.logError(err)
	}
}

func (s *Server) logError(err error) {
	// TODO...
}

func (s *Server) reportErr(err error) {
	// TODO ...
}