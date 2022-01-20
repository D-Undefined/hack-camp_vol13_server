package persistance

import (
	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type threadRepository struct {
	sh SqlHandler
}

func NewThreadRepository(sh SqlHandler) repository.ThreadRepository{
	return &threadRepository{sh:sh}
}

// Thread作成
func (tR threadRepository) CreateThread(thread *model.Thread) error {
	db := tR.sh.db
	return db.Save(&thread).Error
}

// Thread削除
func (tR threadRepository) DeleteThread(thread *model.Thread)error{
	db := tR.sh.db
	return db.Delete(&thread).Error
}

// Thread更新
func (tR threadRepository) UpdateThread(thread *model.Thread) error {
	db := tR.sh.db
	return db.Model(&model.Thread{Id:thread.Id}).Update(thread).Error
}


// IDで Threadを検索
func (tR threadRepository) FindThreadById(id int) (*model.Thread,error){
	db := tR.sh.db
	thread := &model.Thread{Id:id}
	err := db.Find(thread).Error
	if err!=nil{
		return nil,err
	}
	return thread,nil
}


// 全ての Thread を取得
func (tR threadRepository) FindAllThread()([]*model.Thread,error){
	db := tR.sh.db
	threads := []*model.Thread{}
	err := db.Find(&threads).Error
	if err!=nil{
		return nil,err
	}
	return threads,nil
}