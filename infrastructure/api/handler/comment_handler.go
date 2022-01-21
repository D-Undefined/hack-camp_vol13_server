package handler

import (
	"net/http"
	"strconv"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	cR repository.CommentRepository
}

type CommentHandler interface {
	CreateComment(*gin.Context)
	DeleteComment(*gin.Context)
}

func NewCommentHandler(cR repository.CommentRepository) CommentHandler {
	return &commentHandler{cR: cR}
}

// comment作成
func (cH *commentHandler) CreateComment(ctx *gin.Context) {
	comment := &model.Comment{}

	if err := ctx.Bind(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := cH.cR.CreateComment(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)

}

// comment削除
func (cH *commentHandler) DeleteComment(ctx *gin.Context) {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	comment := &model.Comment{Id: id}

	if err := cH.cR.DeleteComment(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})

}
