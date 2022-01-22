package handler

import (
	"net/http"

	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
	"github.com/gin-gonic/gin"
)

type statisticsHandler struct {
	sR repository.StatisticsRepository
}

type StatisticsHandler interface {
	GetStatistics(*gin.Context)
}

func NewStatisticsHandler(sR repository.StatisticsRepository) StatisticsHandler {
	return &statisticsHandler{sR: sR}
}

func (sH *statisticsHandler) GetStatistics(ctx *gin.Context) {
	statistics := sH.sR.GetStatistics()
	ctx.JSON(http.StatusOK, statistics)
}
