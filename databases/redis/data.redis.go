package redis_db

import (
	"context"
	"fmt"
	"log"
	"tes/models"
	"time"
)

func (uc *redisDB) WriteRedis(req models.RedisReq)error{
	// store data using SET command
	err := uc.redis.Set(context.Background(), req.Key, req.Data, time.Duration(req.Exp)*time.Second).Err()
	if err != nil {
		fmt.Printf("unable to SET data. error: %v", err)
		return err
	}
	log.Println("set operation success")

	// get data
	// res,err := rdb.Get(context.Background(), key).Result()
	// if err != nil {
	// 	fmt.Printf("unable to GET data. error: %v", err)
	// 	return err
	// }
	// log.Println("get operation success. result:", res)

	return nil
}

func (uc *redisDB) ReadRedis(req models.RedisReq)error{

	// get data
	res,err := uc.redis.Get(context.Background(), req.Key).Result()
	if err != nil {
		fmt.Printf("unable to GET data. error: %v", err)
		return err
	}
	log.Println("get operation success. result:", res)

	return nil
}