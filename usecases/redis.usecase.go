package usecase

import (
	"tes/models"

	"github.com/sirupsen/logrus"
)

func (uc *usecase) WriteRedis(req models.RedisReq)error{
	// store data using SET command
	// err := uc.redis.Set(context.Background(), req.Key, req.Data, time.Duration(10)*time.Second).Err()
	// log.Println("set operation success")
	err:=uc.redis.WriteRedis(req)
	if err != nil {
		logrus.Error("unable to SET data. error: %v", err)
		return err
	}

	return nil
}