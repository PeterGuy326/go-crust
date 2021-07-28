package logic

import (
	"crust/config"
	"crust/model"
)

// 用户登陆
func Login(input model.LoginInput) (res interface{}, err error) {
	var user model.User
	query := config.DB.Table("users").Select("id").
		Where("name = ? and password = ?", input.Name, input.Password).
		First(&user)

	err = query.Error
	if err != nil {
		return res, err
	}

	return user.Id, nil
}

// 用户注册
func Register(input model.RegisterInput) (err error) {
	user := &model.User{Name: input.Name, Password: input.Password}
	err = config.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
