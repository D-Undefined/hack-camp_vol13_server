package handler

import (
	"net/http"
	"strconv"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type threadHandler struct {
	tR repository.ThreadRepository
}

type ThreadHandler interface{
	CreateThread(ctx *gin.Context)
	DeleteThread(ctx *gin.Context)
	UpdateThread(ctx *gin.Context)
	FindThreadById(ctx *gin.Context)
	FindAllThread(ctx *gin.Context)

}

func NewThreadHandler(tR repository.ThreadRepository) ThreadHandler{
	return &threadHandler{tR:tR}
}


// Thread作成
func (tH threadHandler) CreateThread(ctx *gin.Context){
	thread := &model.Thread{Vote:0}

	if err:=ctx.Bind(thread);err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	if err:=tH.tR.CreateThread(thread);err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,thread)
}


// Thread削除
func (tH threadHandler) DeleteThread(ctx *gin.Context){
	idString := ctx.Param("id")
	id,err := strconv.Atoi(idString)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	if err:=tH.tR.DeleteThread(&model.Thread{Id:id});err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"message":"ok"})

}


// Thread更新
func (tH threadHandler) UpdateThread(ctx *gin.Context){
	idString := ctx.Param("id")
	id,err := strconv.Atoi(idString)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	thread := &model.Thread{Id:id}

	if err:=ctx.Bind(thread);err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	if err:=tH.tR.UpdateThread(thread);err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,thread)
}



// IDで Threadを検索
func (tH threadHandler) FindThreadById(ctx *gin.Context){
	idString := ctx.Param("id")
	id,err := strconv.Atoi(idString)

	if err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	thread,err := tH.tR.FindThreadById(id)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,thread)

}


// 全ての Thread を取得
func (tH threadHandler) FindAllThread(ctx *gin.Context){
	threads,err := tH.tR.FindAllThread()
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,threads)

}