package helper

import (
	"github.com/go-redis/redis/v8"
)

const (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 0
)

var (
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
