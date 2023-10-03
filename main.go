package main

import (
	controller "tes/controllers"
	postgre "tes/databases/postgresql"
	redis_db "tes/databases/redis"
	router "tes/routers"
	usecase "tes/usecases"
)

func main() {
	postgre := postgre.InitPostgre()
	redis := redis_db.InitRedis()
	uc := usecase.InitUsecase(postgre, redis)
	con := controller.InitController(uc)

	router.MainRouter(con)

}

