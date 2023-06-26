package usecase

import (
	"tes/models"

	postgre "tes/databases/postgresql"
	redis_db "tes/databases/redis"
)

type (
	usecase struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
	}
	UsecaseInterface interface {
		WriteRedis(models.RedisReq) error
		ReadRedis(req models.RedisReq) (string, error)
		InsertDB(req models.ItemList) error
	}
)

func InitUsecase(postgre postgre.PostgreInterface, redis redis_db.RedisInterface) UsecaseInterface {
	return &usecase{
		postgre: postgre,
		redis:   redis,
	}
}
