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



// user作成
func (uR *userRepository) CreateUser(user *model.User) error {
	db := uR.sh.db
	return db.Save(&user).Error
}



// uid で userを検索
func (uR *userRepository) FindUserById(uid string) (*model.User, error) {
	db := uR.sh.db

	user := &model.User{Id: uid}
	if err := db.Find(user).Error;err != nil {
		return nil, err
	}

	threads := []*model.Thread{}
	if err := db.Find(&threads).Error; err != nil {
		return nil, err
	}
	user.Threads = threads
	return user, nil
}


// すべてのuserを返す
func (uR *userRepository) FindAllUser() ([]*model.User, error) {
	db := uR.sh.db
	users := []*model.User{}
	err := db.Preload("Threads").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}