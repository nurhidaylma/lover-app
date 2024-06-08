package model

import (
	"time"

	"github.com/nurhidaylma/lover-app.git/util"
)

type User struct {
	Id            int
	Email         string
	UserName      string
	Password      string
	PremiumStatus util.PremiumStatusTypes
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Profile struct {
	Id        int
	UserId    int
	Name      string
	Age       string
	Gender    util.GenderTypes
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Swipe struct {
	Id        int
	UserId    int // the user performing swipe
	ProfileId int // the user being swiped
	SwipeType util.SwipeTypes
	CreatedAt time.Time
}

type PremiumFeature struct {
	Id          int
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
}

type UserPurchase struct {
	Id           int
	UserId       int
	FeatureId    int
	PurchaseDate time.Time
}

type ResponseMessage struct {
	Message string
}
