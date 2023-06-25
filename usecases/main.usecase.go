package usecase

import (
	"tes/models"

	redis_db "tes/databases/redis"
)

type (
	usecase struct{
		redis redis_db.RedisInterface
	}
	UsecaseInterface interface{
		WriteRedis(models.RedisReq)error
	}
)

func InitUsecase(redis redis_db.RedisInterface) UsecaseInterface{
	return &usecase{
		redis:redis,
	}
}
