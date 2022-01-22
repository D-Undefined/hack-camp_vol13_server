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

	// api var.1.0.0
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
	// v1.GET("/thread/user_ranking", tH.UserOfThreadRanking)

	// thread vote
	v1.POST("/vote_thread", vtH.IncreaseVoteThread)
	v1.DELETE("/vote_thread", vtH.RevokeVoteThread)

	// uid が thread_id に対して　投票してるか
	// 投票してた場合 model.VoteThreadの型で返す
	v1.GET("/vote_thread/:uid/:thread_id", vtH.CheckVoteThread)

	// comment
	v1.POST("/comment", cH.CreateComment)
	v1.DELETE("/comment/:id", cH.DeleteComment)

	// comment vote
	v1.POST("/vote_comment", vcH.IncreaseVoteComment)
	v1.DELETE("/vote_comment", vcH.RevokeVoteComment)

	// thread id でコメント一覧を取得し
	// その中でどのコメントにuidが投票したcomment一覧をリストで返す
	// 投票したcommentがないときは空配列を返す
	v1.GET("/vote_comment/:uid/:thread_id", vcH.FindVoteCommentIdOfVoted)

	server.Run(":8080")
}

func health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
