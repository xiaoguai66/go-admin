package repository

import (
	"admin-demo/model"
	"admin-demo/service/request"
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

func (m *UserRepository) AddUser(option *request.UserAddDto) error {
	var userModel = model.User{}
	option.ConvertToModel(&userModel)
	err := m.Orm.Save(&userModel).Error
	if err == nil {
		option.ID = userModel.ID
		option.Password = ""
	}
	return err
}

func (m *UserRepository) CheckUserNameExist(name string) bool {
	var total int64
	m.Orm.Model(model.User{}).Where("name=?", name).Count(&total)
	return total > 0
}
