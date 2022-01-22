package persistance

import (
	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type statisticsRepository struct {
	sh *SqlHandler
}

func NewStatisticsRepository(sh *SqlHandler)repository.StatisticsRepository{
	return &statisticsRepository{sh:sh}
}

func (sR *statisticsRepository) GetStatistics()*model.Statistics{
	db := sR.sh.db
	
	// エラーハンドリング　未実装
	thread_cnt := db.Find(&model.Thread{}).RowsAffected
	comment_cnt := db.Find(&model.Comment{}).RowsAffected
	vote_cnt := db.Find(&model.VoteComment{}).RowsAffected + db.Find(&model.VoteThread{}).RowsAffected

	statistics := &model.Statistics{
		SumThread: int(thread_cnt),
		SumComment: int(comment_cnt),
		SumVote: int(vote_cnt),
	}
	return statistics
}