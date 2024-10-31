package repository

import (
	"admin-demo/model"
)

var (
	userRepository *UserRepository
)

type UserRepository struct {
	*BaseRepository
}

func NewUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{NewBaseRepository()}
	}
	return userRepository
}

func (m *UserRepository) GetUserByNameAndPassword(name, password string) model.User {
	var userModel = model.User{}
	//m.Orm.Table("sys_user").Where("name=? and password=?", name, password).Find(&userModel)
	m.Orm.Model(userModel).Where("name=? and password=?", name, password).Find(&userModel)
	return userModel
}
