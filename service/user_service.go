package service

import (
	"admin-demo/model"
	"admin-demo/repository"
	"admin-demo/service/request"
	"errors"
)

var (
	userService *UserService
)

type UserService struct {
	BaseService
	repository *repository.UserRepository
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			repository: repository.NewUserRepository(),
		}
	}
	return userService
}

func (u *UserService) Login(option request.UserLoginRequest) (model.User, error) {
	var errResult error
	user := u.repository.GetUserByNameAndPassword(option.Name, option.Password)
	if user.ID == 0 {
		errResult = errors.New("用户名或密码错误")
	}
	return user, errResult
}

func (u *UserService) AddUser(option *request.UserAddRequest) error {
	//检查用户名
	if u.repository.CheckUserNameExist(option.Name) {
		return errors.New("用户名已存在")
	}
	return u.repository.AddUser(option)
}

func (u *UserService) GetUserInfoById(idRequest *request.CommonIDRequest) (model.User, error) {
	return u.repository.GetUserInfoById(idRequest.ID)
}

func (u *UserService) GetUserList(option *request.UserListRequest) ([]model.User, int64, error) {
	return u.repository.GetUserList(option)
}

func (u *UserService) UpdateUser(option *request.UserUpdateRequest) error {
	//重名判断
	return u.repository.UpdateUser(option)
}

func (u *UserService) DeleteUserById(option *request.CommonIDRequest) error {
	return u.repository.DeleteUserById(option.ID)
}
