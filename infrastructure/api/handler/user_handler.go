package handler

import (
	"net/http"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uR repository.UserRepository
}

type UserHandler interface {
	FindAllUser(*gin.Context)
	CreateUser(*gin.Context)
	FindUserById(*gin.Context)
}



func NewUserHandler(uR repository.UserRepository) UserHandler {
	return &userHandler{uR: uR}
}


// すべてのuserを返す
func (uH *userHandler) FindAllUser(ctx *gin.Context) {
	users, err := uH.uR.FindAllUser()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}


// user作成
func (uH *userHandler) CreateUser(ctx *gin.Context) {

	// userの初期値設定
	user := &model.User{
		Comment:  "",
		Location: "",
		Twitter:  "",
		Github:   "",
		Url:      "",
		Follow:   0,
		Follower: 0,
	}
	if err := ctx.BindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := uH.uR.CreateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}


// uid で userを検索
func (uH *userHandler) FindUserById(ctx *gin.Context) {
	uid := ctx.Param("uid")
	user, err := uH.uR.FindUserById(uid)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)

}
