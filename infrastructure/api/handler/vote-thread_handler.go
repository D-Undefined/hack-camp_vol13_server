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
	IncreaseThreadVote(*gin.Context)
	RevokeThreadVote(*gin.Context)
}

func NewVoteThreadHandler(vtR repository.VoteThreadRepository) VoteThreadHandler {
	return &voteThreadHandler{vtR: vtR}
}

// good/bad を増やす
func (vtH *voteThreadHandler) IncreaseThreadVote(ctx *gin.Context) {
	thread_vote := &model.ThreadVote{}
	if err := ctx.Bind(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vtH.vtR.IncreaseThreadVote(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// good/bad の取り消し
func (vtH *voteThreadHandler) RevokeThreadVote(ctx *gin.Context) {
	thread_vote := &model.ThreadVote{}
	if err := ctx.Bind(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vtH.vtR.RevokeThreadVote(thread_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
