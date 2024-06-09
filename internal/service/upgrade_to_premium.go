package service

import (
	"context"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LoverService) UpgradeToPremium(ctx context.Context, req model.UserPurchase) (model.ResponseMessage, error) {
	userId := util.GetCtxString(ctx, util.CtxKeyUserId)
	req.UserId = util.StrToInt(userId)

	err := s.repo.DB.WriteUserPurchase(req)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed purchasing to premium"}, status.Error(codes.Internal, codes.Internal.String())
	}

	err = s.repo.DB.UpdateUserPremiumStatus(util.IsPremium, req.UserId)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed purchasing to premium"}, status.Error(codes.Internal, codes.Internal.String())
	}

	go s.repo.Redis.SetPremiumUser(req)
	return model.ResponseMessage{Message: "success"}, nil
}
