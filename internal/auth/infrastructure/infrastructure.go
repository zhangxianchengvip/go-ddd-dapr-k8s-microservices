package infrastructure

import (
	"github.com/redis/go-redis/v9"
	"github.com/zhangxianchengvip/go-ddd-dapr-k8s-microservices/pkg/cache"
)

type Infrastructure struct {
	redis_conn  string
	redis_index int
	redis_pwd   string
	Redis       *redis.Client
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

func (i *Infrastructure) build() (*Infrastructure, error) {

	if redis, err := cache.NewRedisCache("localhost:6379", "", 0); err == nil {
		i.Redis = redis
	} else {
		return nil, err
	}

	return i, nil
}
