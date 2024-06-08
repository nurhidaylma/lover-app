package util

type contextKey string

const (
	CtxKeyUserId = contextKey("UserId")
)

type RedisField string

const (
	RedisFieldUserId       RedisField = "user_id"
	RedisFieldFeatureId    RedisField = "feature_id"
	RedisFieldPurchaseDate RedisField = "purchase_date"
)

type GenderTypes int

const (
	Male    GenderTypes = 1
	Female  GenderTypes = 2
	NotBoth GenderTypes = 3
)

type SwipeTypes int

const (
	Like SwipeTypes = 1
	Pass SwipeTypes = -1
)

type PremiumStatusTypes int

const (
	IsPremium    = 1
	IsNotPremium = -1
)
