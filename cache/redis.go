package cache

import (
	"context"
	"gin-ranking/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisHost,
		Password: "",
		DB:       0,
	})

	Rctx = context.Background()
}

func Zscore(id int, score int) redis.Z {
	return redis.Z{
		Score:  float64(score),
		Member: id,
	}
}
