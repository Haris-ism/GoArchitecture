package router

import (
	controller "tes/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func MainRouter(con controller.ControllerInterface){
	r := gin.Default()
	v1:=r.Group("v1")
	v1.GET("/ping",con.Ping)
	v1.POST("/writeredis",con.WriteRedis)

	logrus.Info("starts")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

