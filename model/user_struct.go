package model

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 用户登陆
type LoginInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginOutput struct {
	Id int `json:"id"`
}

// 用户注册
type RegisterInput struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
