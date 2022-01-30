package helper

import (
	"github.com/go-redis/redis/v8"
)

const (
	addr     = "localhost:6379"
	password = ""
	db       = 0
)

var (
	rdb *redis.Client
)

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func GetRedis() *redis.Client {
	return rdb
}
