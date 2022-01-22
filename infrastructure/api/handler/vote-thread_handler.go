package handler

import (
	"net/http"
	"strconv"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type voteThreadHandler struct {
	vtR repository.VoteThreadRepository
}

type VoteThreadHandler interface {
	IncreaseVoteThread(*gin.Context)
	RevokeVoteThread(*gin.Context)
	CheckVoteThread(*gin.Context)
}

func NewVoteThreadHandler(vtR repository.VoteThreadRepository) VoteThreadHandler {
	return &voteThreadHandler{vtR: vtR}
}

// good/bad を増やす
func (vtH *voteThreadHandler) IncreaseVoteThread(ctx *gin.Context) {
	thread_vote := &model.VoteThread{}
	if err := ctx.Bind(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vtH.vtR.IncreaseVoteThread(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// good/bad の取り消し
func (vtH *voteThreadHandler) RevokeVoteThread(ctx *gin.Context) {
	thread_vote := &model.VoteThread{}
	if err := ctx.Bind(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vtH.vtR.RevokeVoteThread(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// good/bad 済みか
func (vtH *voteThreadHandler) CheckVoteThread(ctx *gin.Context) {
	uid := ctx.Param("uid")

	threadIdString := ctx.Param("thread_id")
	threadId, err := strconv.Atoi(threadIdString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	vote_thread, err := vtH.vtR.CheckVoteThread(uid, threadId)

	/* error でも StatusOKにしてます
	理由は存在しなかった場合でも正常処理を行いたいからです
	ほかにいい書き方があれば絶対に変えたほうが良いです*/
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message":"not found"})
		return
	}

	ctx.JSON(http.StatusOK, vote_thread)
}
