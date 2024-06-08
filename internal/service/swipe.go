package service

import (
	"context"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *LoverService) Swipe(ctx context.Context, req model.Swipe) (model.ResponseMessage, error) {
	const swipeLimit = 10
	userId := util.GetCtxString(ctx, util.CtxKeyUserId)
	req.UserId = util.StrToInt(userId)

	if !util.ValidSwipeType(int(req.SwipeType)) {
		util.Logger.LogWarning("swipe is invalid")
		return model.ResponseMessage{Message: "swipe is invalid"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	currentCount, err := s.repo.Redis.GetSwipeCount(req)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed swiping"}, status.Error(codes.Internal, codes.Internal.String())
	}

	isPremium, err := s.repo.Redis.GetPremiumUser(model.UserPurchase{
		UserId: req.UserId,
	})
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed swiping"}, status.Error(codes.Internal, codes.Internal.String())
	}

	if !isPremium {
		if currentCount > swipeLimit {
			util.Logger.LogWarning("swipe reached limit")
			return model.ResponseMessage{Message: "swipe reached limit"}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
		}
	}

	err = s.repo.DB.WriteSwipe(req)
	if err != nil {
		util.Logger.LogError(err.Error())
		return model.ResponseMessage{Message: "failed writing profile"}, status.Error(codes.Internal, codes.Internal.String())
	}

	go s.repo.Redis.SetSwipeCount(req)

	return model.ResponseMessage{Message: "success"}, nil
}
