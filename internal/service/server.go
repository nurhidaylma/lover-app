package service

import (
	"context"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/internal/repository"
)

type LoverService struct {
	repo repository.Repository
}

type Service interface {
	Login(email, password string) (model.User, error)
	SignUp(model.User) (model.ResponseMessage, error)
	SetProfile(context.Context, model.Profile) (model.ResponseMessage, error)

	Swipe(context.Context, model.Swipe) (model.ResponseMessage, error)

	UpgradeToPremium(context.Context, model.UserPurchase) (model.ResponseMessage, error)
}

func NewLoverService(repo repository.Repository) *LoverService {
	return &LoverService{
		repo: repo,
	}
}
