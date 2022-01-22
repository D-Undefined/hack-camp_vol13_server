package persistance

import (
	"fmt"

	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type userRepository struct {
	sh *SqlHandler
}

func NewUserRepository(sh *SqlHandler) repository.UserRepository {
	return &userRepository{sh: sh}
}

// user作成
func (uR *userRepository) CreateUser(user *model.User) (*model.User, error) {
	db := uR.sh.db

	// uidがあるかどうか
	if user.Id == "" {
		return nil, fmt.Errorf("uid is empty")
	}

	//存在するか確認
	// 既に存在する場合　resUserにデータをバインドして返す
	resUser := &model.User{Id: user.Id}
	if err := db.First(resUser).Error; err == nil {
		return resUser, nil
	}

	return user, db.Save(user).Error
}

// user 削除
func (uR *userRepository) DeleteUser(user *model.User) error {
	db := uR.sh.db
	//存在するか確認
	if err := db.First(&model.User{Id: user.Id}).Error; err != nil {
		return err
	}

	return db.Delete(user).Error
}

// user 更新
func (uR *userRepository) UpdateUser(user *model.User) error {
	db := uR.sh.db
	//存在するか確認
	if err := db.First(&model.User{Id: user.Id}).Error; err != nil {
		return err
	}

	return db.Model(&model.User{Id: user.Id}).Update(user).Error
}

// uid で userを検索
func (uR *userRepository) FindUserById(uid string) (*model.User, error) {
	db := uR.sh.db

	user := &model.User{Id: uid}
	err := db.Preload("Threads").First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// すべてのuserを返す(threadまで)
func (uR *userRepository) FindAllUser() (*[]*model.User, error) {
	db := uR.sh.db

	users := &[]*model.User{
		{
			Threads: []*model.Thread{},
		},
	}

	err := db.Preload("Threads").Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// 上位10名の userを表示 (Scoreを基準とする)
func (uR *userRepository) GetUserRanking() (*[]*model.User, error) {
	db := uR.sh.db
	top_users := &[]*model.User{}
	if err := db.Limit(2).Order("score desc").Find(top_users).Error; err != nil {
		return nil, err
	}
	return top_users, nil
}

// [おまけ] すべてのuserを返す(commentまで結合)
// func (uR *userRepository) FindAllUser() (*[]*model.User, error) {
// 	db := uR.sh.db

// 	users := &[]*model.User{
// 		{
// 			Threads: []*model.Thread{
// 				{
// 					Comments: []*model.Comment{},
// 				},
// 			},
// 		},
// 	}

// 	/*
// 		返り値の例
// 		[
// 			{
// 				"uid": "a38ty89haeh",
// 				"user_name": "hoge",...
// 				"Threads": [
// 					{
// 						"id":1,....
// 						"Comments":[{},{},{},...]
// 					}
// 				]

// 			},
// 			{},
// 			{},...
// 		]
// 	*/

// 	err := db.Preload("Threads.Comments").Find(users).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
