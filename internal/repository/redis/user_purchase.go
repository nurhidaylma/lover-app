package redis

import (
	"fmt"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"
)

func (r *redisClient) SetPremiumUser(req model.UserPurchase) error {
	const (
		fileName = `premium_feature.go`
		funcName = `SetPremiumUser`
	)

	key := fmt.Sprintf("PREMIUM:%d", req.UserId)
	data := make(map[string]interface{})
	data[string(util.RedisFieldFeatureId)] = req.FeatureId
	data[string(util.RedisFieldPurchaseDate)] = req.PurchaseDate

	_, err := r.client.HMSet(key, data).Result()
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}

	return err
}

func (r *redisClient) GetPremiumUser(req model.UserPurchase) (bool, error) {
	const (
		fileName = `premium_feature.go`
		funcName = `GetPremiumUser`
	)

	key := fmt.Sprintf("PREMIUM:%d", req.UserId)
	value, err := r.client.HGetAll(key).Result()
	if err != nil {
		return false, util.NewError(fileName, funcName, err)
	}

	if len(value) == 0 {
		return false, nil
	}

	return true, nil
}
