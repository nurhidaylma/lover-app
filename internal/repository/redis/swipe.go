package redis

import (
	"fmt"
	"time"

	"github.com/nurhidaylma/lover-app.git/internal/model"
	"github.com/nurhidaylma/lover-app.git/util"

	"github.com/go-redis/redis/v7"
)

func (r *redisClient) SetSwipeCount(req model.Swipe) error {
	const (
		fileName = `swipe.go`
		funcName = `SetSwipeCount`
	)

	key := fmt.Sprintf("SWIPES:%d", req.UserId)
	count, err := r.client.Incr(key).Result()
	if err != nil {
		return util.NewError(fileName, funcName, err)
	}

	if count == 1 {
		r.client.Expire(key, time.Hour*24)
	}

	return nil
}

func (r *redisClient) GetSwipeCount(req model.Swipe) (int, error) {
	const (
		fileName = `swipe.go`
		funcName = `GetSwipeCount`
	)

	key := fmt.Sprintf("SWIPES:%d", req.UserId)
	count, err := r.client.Get(key).Int()
	if err != nil && err != redis.Nil {
		return 0, util.NewError(fileName, funcName, err)
	}

	return count, nil
}
