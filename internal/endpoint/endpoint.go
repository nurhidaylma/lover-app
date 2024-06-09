package endpoint

import (
	"context"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/internal/service"
)

type LoverEndpoint struct {
	SignUpEndpoint           func(user model.User) (model.ResponseMessage, error)
	LoginEndpoint            func(email, password string) (model.User, error)
	SetProfileEndpoint       func(context.Context, model.Profile) (model.ResponseMessage, error)
	SwipeEndpoint            func(context.Context, model.Swipe) (model.ResponseMessage, error)
	UpgradeToPremiumEndpoint func(context.Context, model.UserPurchase) (model.ResponseMessage, error)
}

func MakeLoverEndpoints(s *service.LoverService) LoverEndpoint {
	signUpEndpoint := func(user model.User) (model.ResponseMessage, error) {
		return s.SignUp(user)
	}

	loginEndpoint := func(email, password string) (model.User, error) {
		return s.Login(email, password)
	}

	setProfileEndpoint := func(ctx context.Context, req model.Profile) (model.ResponseMessage, error) {
		return s.SetProfile(ctx, req)
	}

	swipeEndpoint := func(ctx context.Context, req model.Swipe) (model.ResponseMessage, error) {
		return s.Swipe(ctx, req)
	}

	upgradeToPremiumEndpoint := func(ctx context.Context, req model.UserPurchase) (model.ResponseMessage, error) {
		return s.UpgradeToPremium(ctx, req)
	}

	return LoverEndpoint{
		SignUpEndpoint:           signUpEndpoint,
		LoginEndpoint:            loginEndpoint,
		SetProfileEndpoint:       setProfileEndpoint,
		SwipeEndpoint:            swipeEndpoint,
		UpgradeToPremiumEndpoint: upgradeToPremiumEndpoint,
	}
}
