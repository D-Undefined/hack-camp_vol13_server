package persistance

import (
	"fmt"
	"time"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type threadRepository struct {
	sh *SqlHandler
}

func NewThreadRepository(sh *SqlHandler) repository.ThreadRepository {
	return &threadRepository{sh: sh}
}

// Thread作成
func (tR *threadRepository) CreateThread(thread *model.Thread) (*model.Thread, error) {
	db := tR.sh.db

	// uidがあるかどうか
	if thread.UserID == "" {
		return nil, fmt.Errorf("uid is empty")
	}

	// userが存在するか確認
	user := &model.User{Id: thread.UserID}
	if err := db.First(user).Error; err != nil {

		return nil, err
	}

	//作成した user の　scoreを30増やす
	user.Score = user.Score + 30
	if err := db.Model(&model.User{Id: thread.UserID}).Update(user).Error; err != nil {
		return nil, err
	}

	// thread Save
	if err := db.Save(thread).Error; err != nil {
		return nil, err
	}

	// Saveすることでthreadに id がふられる
	// それを利用してresponse用のThreadデータ準備
	resThread := &model.Thread{
		Id:   thread.Id,
		User: &model.User{},
	}
	if err := db.Preload("User").First(resThread).Error; err != nil {
		return nil, err
	}

	return resThread, nil
}

// Thread削除
func (tR *threadRepository) DeleteThread(thread *model.Thread) error {
	db := tR.sh.db
	//存在するか確認
	if err := db.First(&model.Thread{Id: thread.Id}).Error; err != nil {
		return err
	}

	// CASCADE 実装したかった (できてない)
	// thread に　Commentを結合
	// if err:=db.Model(&model.Thread{Id: thread.Id}).Association("Comments").Error;err!=nil{
	// 	return err
	// }

	// if err:=db.Preload("Comments").Find(thread).Error;err!=nil{
	// 	return err
	// }

	// return db.Select("Comments").Delete(thread).Error

	return db.Delete(thread).Error

}

// Thread更新
func (tR *threadRepository) UpdateThread(thread *model.Thread) (*model.Thread, error) {
	db := tR.sh.db
	//存在するか確認
	if err := db.First(&model.Thread{Id: thread.Id}).Error; err != nil {
		return nil, err
	}

	// 更新
	if err := db.Model(&model.Thread{Id: thread.Id}).Update(thread).Error; err != nil {
		return nil, err
	}

	// response用のThreadデータ準備
	resThread := &model.Thread{
		Id:   thread.Id,
		User: &model.User{},
	}
	if err := db.Preload("User").First(resThread).Error; err != nil {
		return nil, err
	}

	return resThread, nil
}

// IDで Threadを検索
func (tR *threadRepository) FindThreadById(id int) (*model.Thread, error) {
	db := tR.sh.db
	thread := &model.Thread{
		Id: id,
		User: &model.User{},
		Comments: []*model.Comment{
			{
				User: &model.User{},
			},
		},
	}

	err := db.Preload("Comments.User").First(thread).Error
	if err != nil {
		return nil, err
	}
	return thread, nil
}

// 全ての Thread を取得
func (tR *threadRepository) FindAllThread() (*[]*model.Thread, error) {
	db := tR.sh.db
	threads := &[]*model.Thread{
		{
			User: &model.User{},
			Comments: []*model.Comment{},
		},
	}
	err := db.Preload("User").Preload("Comments").Find(threads).Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}

// 過去１週間の trendのthreadを 10件返す
func (tR *threadRepository) FindTrendThread() (*[]*model.Thread, error) {
	var now, lastWeek time.Time

	db := tR.sh.db
	trend_thread := &[]*model.Thread{}

	// test してないので 正しく動くか不安
	now = time.Now()
	lastWeek = now.AddDate(0, 0, -7)

	if err := db.Where("created_at BETWEEN ? AND ?", lastWeek, now).
		Limit(10).
		Order("vote_cnt desc").
		Find(trend_thread).Error; err != nil {
		return nil, err
	}

	return trend_thread, nil
}

// Thread(VoteCnt)の user ランキング
// Threadが0のユーザーはランキングに乗らない
// func (tR *threadRepository) UserOfThreadRanking() (*[]*model.UserRanking, error) {
// 	db := tR.sh.db

// 	results := &[]*model.UserRanking{
// 		{
// 			User: &model.User{},
// 		},
// 	}

// 	if err := db.Table("threads").
// 				Select("user_id, sum(vote_cnt) as vote_sum").
// 				Group("user_id").
// 				Limit(50).
// 				Find(results).Error;
// 				err != nil {
// 					return nil, err
// 	}

// 	fmt.Printf("### %v\n",results)

// 	if err:=db.Joins("left join users on users.id = user_rankings.user_id").
// 				// Order("vote_sum").
// 				Scan(results).Error;
// 				err!=nil{
// 					return nil,err
// 	}

// 	fmt.Printf("$$$ %v\n",results)

// 	return results, nil
// }
