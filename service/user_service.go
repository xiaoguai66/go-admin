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

func (u *UserService) Login(option request.UserLoginDto) (model.User, error) {
	var errResult error
	user := u.repository.GetUserByNameAndPassword(option.Name, option.Password)
	if user.ID == 0 {
		errResult = errors.New("用户名或密码错误")
	}
	return user, errResult
}
