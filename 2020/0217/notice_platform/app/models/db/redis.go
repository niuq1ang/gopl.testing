package db

import (
	"github.com/bangwork/ones-plugin/notice_platform/app/utils/config"
	"gopkg.in/redis.v5"
)

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisAddr,
		Password: config.Config.RedisPassword, // no password set
		DB:       config.Config.RedisDB,       // use default DB
	})
	return RedisClient.FlushDb().Err()
}
