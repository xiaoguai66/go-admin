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

func (m *UserRepository) AddUser(option *request.UserAddRequest) error {
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

func (m *UserRepository) GetUserInfoById(id int32) (model.User, error) {
	var userModel model.User
	err := m.Orm.First(&userModel, id).Error
	return userModel, err
}

func (m *UserRepository) GetUserList(option *request.UserListRequest) ([]model.User, int64, error) {
	var userList []model.User
	var total int64
	err := m.Orm.Model(&model.User{}).
		Scopes(Paginage(option.Paginate)).Find(&userList).
		Offset(-1).Limit(-1).Count(&total).Error
	return userList, total, err
}

func (m *UserRepository) UpdateUser(option *request.UserUpdateRequest) error {
	var user model.User
	m.Orm.First(&user, option.ID)
	option.ConvertToModel(&user)

	return m.Orm.Save(&user).Error
}

func (m *UserRepository) DeleteUserById(id int32) error {
	return m.Orm.Delete(&model.User{}, id).Error
}
