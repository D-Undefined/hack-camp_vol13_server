package repository

import "github.com/D-Undefined/hack-camp_vol13_server/domain/model"

type StatisticsRepository interface {
	GetStatistics() *model.Statistics
}