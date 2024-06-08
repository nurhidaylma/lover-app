package endpoint

import (
	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/internal/service"
)

type LoverEndpoint struct {
	SignUpEndpoint func(user model.User) (model.ResponseMessage, error)
	LoginEndpoint  func(email, password string) (model.User, error)
}

func MakeLoverEndpoints(s *service.LoverService) LoverEndpoint {
	signUpEndpoint := func(user model.User) (model.ResponseMessage, error) {
		return s.SignUp(user)
	}

	loginEndpoint := func(email, password string) (model.User, error) {
		return s.Login(email, password)
	}

	return LoverEndpoint{
		SignUpEndpoint: signUpEndpoint,
		LoginEndpoint:  loginEndpoint,
	}
}
