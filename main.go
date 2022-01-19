package main

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func main(){
	server := gin.Default()

	server.GET("/",health)

	server.Run(":8080")
}

func health(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"message":"ok",
	})
}