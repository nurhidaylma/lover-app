package service

import (
	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LoverService) Login(email, password string) (model.User, error) {
	user, err := s.repo.DB.ReadUserByEmail(email)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.User{}, status.Error(codes.Internal, codes.Internal.String())
	}
	if user.Id == 0 {
		util.Logger.LogWarning("user is not registered")
		return model.User{}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		util.Logger.LogError("email or password is invalid")
		return model.User{}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	return user, nil
}
