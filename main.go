package main

import (
	"net/http"
	"os"

	"github.com/D-Undefined/hack-camp_vol13_server/infrastructure/api/handler"
	"github.com/D-Undefined/hack-camp_vol13_server/infrastructure/persistance"
	"github.com/gin-gonic/gin"
)

func main() {
	// db 準備
	sh := persistance.NewDB()

	// repository
	uR := persistance.NewUserRepository(sh)
	tR := persistance.NewThreadRepository(sh)
	cR := persistance.NewCommentRepository(sh)
	vcR := persistance.NewVoteCommentRepository(sh)
	vtR := persistance.NewVoteThreadRepository(sh)

	// handler
	uH := handler.NewUserHandler(uR)
	tH := handler.NewThreadHandler(tR)
	cH := handler.NewCommentHandler(cR)
	vcH := handler.NewVoteCommentHandler(vcR)
	vtH := handler.NewVoteThreadHandler(vtR)

	// server 準備
	server := gin.Default()

	// test server
	server.GET("/", health)

	// api var.1
	v1 := server.Group("/api/v1")

	// user
	v1.GET("/users", uH.FindAllUser)
	v1.GET("/user/:uid", uH.FindUserById)
	v1.POST("/user", uH.CreateUser)
	v1.PUT("/user/:uid", uH.UpdateUser)
	v1.DELETE("/user/:uid", uH.DeleteUser)

	// thread
	v1.GET("/threads", tH.FindAllThread)
	v1.GET("/thread/:id", tH.FindThreadById)
	v1.POST("/thread", tH.CreateThread)
	v1.PUT("/thread/:id", tH.UpdateThread)
	v1.DELETE("/thread/:id", tH.DeleteThread)

	// thread vote
	v1.POST("/thread_vote", vtH.IncreaseThreadVote)
	v1.DELETE("/thread_vote", vtH.RevokeThreadVote)

	// comment
	v1.POST("/comment", cH.CreateComment)
	v1.DELETE("/comment/:id", cH.DeleteComment)

	// comment vote
	v1.POST("/comment_vote", vcH.IncreaseCommentVote)
	v1.DELETE("/comment_vote", vcH.RevokeCommentVote)

	server.Run(":" + os.Getenv("PORT"))
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
