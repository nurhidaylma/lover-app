package interfaces

import "github.com/nurhidaylma/lover-app.git/internal/model"

type DBRepository interface {
	WriteUser(user model.User) error
	ReadUserByUserName(username string) (model.User, error)
	ReadUserByEmail(email string) (model.User, error)
	ReadUserById(id int) (model.User, error)
	UpdateUserPremiumStatus(premiumStatus, userID int) error

	WriteProfile(model.Profile) error

	WriteSwipe(model.Swipe) error

	WriteUserPurchase(model.UserPurchase) error
}

type RedisRepository interface {
	SetSwipeCount(model.Swipe) error
	GetSwipeCount(model.Swipe) (int, error)

	SetPremiumUser(model.UserPurchase) error
	GetPremiumUser(model.UserPurchase) (bool, error)
}
