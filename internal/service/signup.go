package service

import (
	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LoverService) SignUp(req model.User) (model.ResponseMessage, error) {
	if !util.ValidEmailAddress(req.Email) {
		util.Logger.LogWarning("invalid email address")
		return model.ResponseMessage{Message: "invalid email address"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	userByEmail, err := s.repo.DB.ReadUserByEmail(req.Email)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed signing up"}, status.Error(codes.Internal, codes.Internal.String())
	}
	if userByEmail.Id != 0 {
		util.Logger.LogWarning("user already registered")
		return model.ResponseMessage{Message: "user already registered"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	userByUsername, err := s.repo.DB.ReadUserByUserName(req.UserName)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed signing up"}, status.Error(codes.Internal, codes.Internal.String())
	}
	if userByUsername.Id != 0 {
		util.Logger.LogWarning("user already registered")
		return model.ResponseMessage{Message: "user already registered"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed signing up"}, status.Error(codes.Internal, codes.Internal.String())
	}
	req.Password = string(hashedPassword)

	err = s.repo.DB.WriteUser(req)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed signing up"}, status.Error(codes.Internal, codes.Internal.String())
	}

	return model.ResponseMessage{Message: "success"}, nil
}
