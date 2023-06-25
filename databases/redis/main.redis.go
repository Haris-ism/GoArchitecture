package redis_db

import (
	"fmt"
	"tes/models"
	"tes/utils"

	"github.com/go-redis/redis/v8"
)

type (
	redisDB struct{
		redis *redis.Client
	}
	RedisInterface interface{
		WriteRedis(models.RedisReq)error
	}
)

func InitRedis() RedisInterface {
	host:=utils.GetEnv("REDIS_HOST")
	password:=utils.GetEnv("REDIS_PASSWORD")
	fmt.Println("ieu cred:",host,password)
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return &redisDB{
		redis:client,
	}
}