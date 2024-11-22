package cache

import (
	"crypto/tls"

	"github.com/redis/go-redis/v9"
)

func NewRedisCache(addr, password string, db int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // 连接地址
		Password: password, // 没有密码，默认值
		DB:       db,       // 默认DB 0
	})

	return rdb
}

func NewRedisCacheByURL(url string) *redis.Client {
	opt, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	return rdb
}

func NewRedisCacheByTLS(serverName string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: serverName,
		},
	})
	return rdb
}
