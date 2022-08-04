package logic

import (
	"crust/dao"
	"crust/model"
)

// 用户登陆
func Login(input model.LoginInput) (res interface{}, err error) {
	user, err := dao.FindUserById(1)
	if err != nil {
		return nil, err
	}
	return user.ID, nil
}

// 用户注册
func Register(input model.RegisterInput) (err error) {
	_, err = dao.AddUser(input.Name, input.Password)
	if err != nil {
		return err
	}
	return nil
}
