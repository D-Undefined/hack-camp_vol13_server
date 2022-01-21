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
func (uR *userRepository) CreateUser(user *model.User) error {
	db := uR.sh.db
	//存在するか確認
	if err := db.First(&model.User{Id: user.Id}).Error; err == nil {
		return fmt.Errorf("this uid already exists")
	}

	return db.Save(user).Error
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

// すべてのuserを返す
func (uR *userRepository) FindAllUser() ([]*model.User, error) {
	db := uR.sh.db
	users := []*model.User{}
	err := db.Preload("Threads").Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
