package request

import "admin-demo/model"

type UserLoginDto struct {
	Name     string `json:"name" binding:"required,first_is_a" message:"用户名错误" required_err:"用户名不能为空" first_is_a_err:"用户名以a字母开头"`
	Password string `json:"password" binding:"required" message:"密码错误" required_err:"密码不能为空"`
}

type UserAddDto struct {
	ID       int32  `json:"id"`
	Name     string `json:"name" form:"name" binding:"required,first_is_a" message:"用户名必填" required_err:"用户名不能为空" first_is_a_err:"用户名以a字母开头"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"密码必填" required_err:"密码不能为空"`
	Mobile   string `json:"mobile" form:"mobile"`
}

func (u *UserAddDto) ConvertToModel(userModel *model.User) {
	userModel.Name = u.Name
	userModel.Mobile = u.Mobile
	userModel.Password = u.Password
}
