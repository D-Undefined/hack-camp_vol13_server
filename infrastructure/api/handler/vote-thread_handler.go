package handler

import (
	"net/http"

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
