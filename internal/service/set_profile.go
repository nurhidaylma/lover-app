package service

import (
	"context"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LoverService) SetProfile(ctx context.Context, req model.Profile) (model.ResponseMessage, error) {
	userId := util.GetCtxString(ctx, util.CtxKeyUserId)
	user, err := s.repo.DB.ReadUserById(util.StrToInt(userId))
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed writing profile"}, status.Error(codes.Internal, codes.Internal.String())
	}
	if user.Id == 0 {
		util.Logger.LogWarning("user is not registered")
		return model.ResponseMessage{Message: "user is not registered"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	if !util.ValidPhoneNumber(req.Phone) {
		util.Logger.LogWarning("invalid phone number")
		return model.ResponseMessage{Message: "invalid phone number"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	req.UserId = user.Id
	err = s.repo.DB.WriteProfile(req)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed writing profile"}, status.Error(codes.Internal, codes.Internal.String())
	}

	return model.ResponseMessage{Message: "success"}, nil
}
