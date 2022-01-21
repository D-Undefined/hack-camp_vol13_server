package persistance

import (
	"fmt"

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
func (tR *threadRepository) CreateThread(thread *model.Thread) error {
	db := tR.sh.db

	// uidがあるかどうか
	if thread.UserID == "" {
		return fmt.Errorf("uid is empty")
	}
	// userが存在するか確認
	if err := db.First(&model.User{Id: thread.UserID}).Error; err != nil {
		return err
	}


	return db.Save(thread).Error
}

// Thread削除
func (tR *threadRepository) DeleteThread(thread *model.Thread) error {
	db := tR.sh.db
	//存在するか確認
	if err := db.First(&model.Thread{Id: thread.Id}).Error; err != nil {
		return err
	}

	// thread に　Commentを結合
	// if err:=db.Model(&model.Thread{Id: thread.Id}).Association("Comments").Error;err!=nil{
	// 	return err
	// }

	// if err:=db.Preload("Comments").Find(thread).Error;err!=nil{
	// 	return err
	// }

	// fmt.Printf("@@@@@@ %v\n",thread)

	// return db.Select("Comments").Delete(thread).Error
	return db.Delete(thread).Error

}

// Thread更新
func (tR *threadRepository) UpdateThread(thread *model.Thread) error {
	db := tR.sh.db
	//存在するか確認
	if err := db.First(&model.Thread{Id: thread.Id}).Error; err != nil {
		return err
	}

	return db.Model(&model.Thread{Id: thread.Id}).Update(thread).Error
}

// IDで Threadを検索
func (tR *threadRepository) FindThreadById(id int) (*model.Thread, error) {
	db := tR.sh.db
	thread := &model.Thread{
		Id:id,
		Comments: []*model.Comment{},
	}

	err := db.Preload("Comments").First(thread).Error
	if err != nil {
		return nil, err
	}
	return thread, nil
}

// 全ての Thread を取得
func (tR threadRepository) FindAllThread() (*[]*model.Thread, error) {
	db := tR.sh.db
	threads := &[]*model.Thread{}
	err := db.Find(threads).Error
	if err != nil {
		return nil, err
	}
	return threads, nil
}


// good に投票
// func (tR threadRepository) VoteGood(id int)error{
// 	db := tR.sh.db
// 	thread := &model.Thread{Id:id}

// 	if err:=db.First(thread).Error;err!=nil{
// 		return err
// 	}
	
// 	thread.VoteGood = thread.VoteGood+1

// }
