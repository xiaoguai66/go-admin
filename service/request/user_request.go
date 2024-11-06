package request

import "admin-demo/model"

type UserLoginRequest struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"用户名错误" required_err:"用户名不能为空" first_is_a_err:"用户名以a字母开头"`
	Password string `json:"password" binding:"required" message:"密码错误" required_err:"密码不能为空"`
}

type UserAddRequest struct {
	ID       int32  `json:"id"`
	Name     string `json:"name" form:"name" binding:"required,first_is_a" message:"用户名必填" required_err:"用户名不能为空" first_is_a_err:"用户名以a字母开头"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码必填" required_err:"密码不能为空"`
	Mobile   string `json:"mobile" form:"mobile"`
}

func (u *UserAddRequest) ConvertToModel(userModel *model.User) {
	userModel.Name = u.Name
	userModel.Mobile = u.Mobile
	userModel.Password = u.Password
}

type UserListRequest struct {
	Paginate
	ID     int32  `json:"id" form:"id" uri:"id"`
	Name   string `json:"name" form:"name" uri:"name"`
	Mobile string `json:"mobile" form:"mobile" uri:"mobile"`
}

type UserUpdateRequest struct {
	ID       int32  `json:"id" uri:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required,first_is_a" message:"用户名必填" required_err:"用户名不能为空" first_is_a_err:"用户名以a字母开头"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码必填" required_err:"密码不能为空"`
	Mobile   string `json:"mobile" form:"mobile"`
}

func (u *UserUpdateRequest) ConvertToModel(userModel *model.User) {
	userModel.ID = u.ID
	userModel.Name = u.Name
	userModel.Mobile = u.Mobile
	userModel.Password = u.Password
}
