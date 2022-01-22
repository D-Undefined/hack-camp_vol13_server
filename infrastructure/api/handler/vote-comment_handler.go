package handler

import (
	"net/http"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type voteCommentHandler struct {
	vcR repository.VoteCommentRepository
}

type VoteCommentHandler interface {
	IncreaseVoteComment(*gin.Context)
	RevokeVoteComment(*gin.Context)
}

func NewVoteCommentHandler(vcR repository.VoteCommentRepository) VoteCommentHandler {
	return &voteCommentHandler{vcR: vcR}
}

// good/bad を増やす
func (vcH *voteCommentHandler) IncreaseVoteComment(ctx *gin.Context) {
	comment_vote := &model.VoteComment{}
	if err := ctx.Bind(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vcH.vcR.IncreaseVoteComment(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// good/bad の取り消し
func (vcH *voteCommentHandler) RevokeVoteComment(ctx *gin.Context) {
	comment_vote := &model.VoteComment{}
	if err := ctx.Bind(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := vcH.vcR.RevokeVoteComment(comment_vote); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
