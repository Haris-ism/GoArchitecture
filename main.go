package main

import (
	controller "tes/controllers"
	redis_db "tes/databases/redis"
	router "tes/routers"
	usecase "tes/usecases"
)

func main() {
	redis:=redis_db.InitRedis()
	uc:=usecase.InitUsecase(redis)
	con:=controller.InitController(uc)
	
	router.MainRouter(con)
	
}


func tesRedis(){
	// var redisHost = "localhost:6379"
	// var redisPassword = ""

	// rdb := newRedisClient(redisHost, redisPassword)
	// fmt.Println("redis client initialized")

	// // // ...

	// key := "key-1"
	// data := "Redis from mac"
	// ttl := time.Duration(100) * time.Second

	// // store data using SET command
	// err := rdb.Set(context.Background(), key, data, ttl).Err()
	// if err != nil {
	// 	fmt.Printf("unable to SET data. error: %v", err)
	// 	return
	// }
	// log.Println("set operation success")

	// // get data
	// res,err := rdb.Get(context.Background(), key).Result()
	// if err != nil {
	// 	fmt.Printf("unable to GET data. error: %v", err)
	// 	return
	// }
	// log.Println("get operation success. result:", res)

	// redisDB:=redis_db.InitRedis()
}