package persistance

import (
	"github.com/D-Undefined/hack-camp_vol13_server/domain/model"
	"github.com/D-Undefined/hack-camp_vol13_server/usecase/repository"
)

type userRepository struct {
	sh SqlHandler
}

func NewUserRepository(sh SqlHandler) repository.UserRepository {
	return &userRepository{sh: sh}
}

// すべてのuserを返す
func (uR *userRepository) FindAllUser() ([]*model.User, error) {
	db := uR.sh.db
	users := []*model.User{}
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}



// user作成
func (uR *userRepository) CreateUser(user *model.User) error {
	db := uR.sh.db
	return db.Save(&user).Error
}



// uid で userを検索
func (uR *userRepository) FindUserById(uid string) (*model.User, error) {
	db := uR.sh.db
	user := &model.User{Uid: uid}
	err := db.Find(user).Error

	if err != nil {
		return nil, err
	}
	return user, err
}
