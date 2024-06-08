package redis

import (
	"github.com/nurhidaylma/lover-app.git/internal/repository"

	"github.com/go-redis/redis/v7"
)

type redisClient struct {
	client *redis.Client
}

type RedisConf struct {
	Address  string
	Password string
}

func (cnf *RedisConf) NewRepository(r *repository.Repository) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cnf.Address,
		Password: cnf.Password,
	})

	r.Redis = &redisClient{client: rdb}
	return nil
}
