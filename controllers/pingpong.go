package controller

import (
	"net/http"
	"tes/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *controller) Ping(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message:"pong euy",
		Code:200,
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *controller) WriteRedis(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message:"success write euy",
		Code:http.StatusOK,
	}
	req:=models.RedisReq{}
	if err:=ctx.BindJSON(&req);err!=nil{
		logrus.Error("ieu error body:",err)
		res.Message="Invalid Body Request"
		res.Code=http.StatusBadRequest
		ctx.JSON(res.Code, res)
		return
	}
	if err:=c.usecase.WriteRedis(req);err!=nil{
		logrus.Error("ieu error write redis:",err)
		res.Message="Failed to Write Redis"
		res.Code=http.StatusInternalServerError
		ctx.JSON(res.Code, res)
		return
	}
	ctx.JSON(http.StatusOK, res)
}