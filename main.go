package main

import (
	"net/http"

	"github.com/D-Undefined/hack-camp_vol13_server/infrastructure/api/handler"
	"github.com/D-Undefined/hack-camp_vol13_server/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

func main() {
	// db 準備
	sh := persistance.NewDB()

	// repository
	uR := persistance.NewUserRepository(*sh)

	// handler
	uH := handler.NewUserHandler(uR)

	// server 準備
	server := gin.Default()

	// test server
	server.GET("/", health)

	// api var.1
	v1 := server.Group("/api/v1")
	v1.GET("/users", uH.FindAllUser)
	v1.POST("/user", uH.CreateUser)
	v1.GET("/user/:uid", uH.FindUserById)

	server.Run(":8080")
}




func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
