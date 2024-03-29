package dao

import (
	"crust/config"
	"time"
)

const TABLE_NAME = "users"

type User struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
}

func AddUser(name string, password string) (res *User, err error) {
	user := &User{Name: name, Password: password}
	err = config.DB.Table(TABLE_NAME).Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserById(id uint) (user User, err error) {
	query := config.DB.Table(TABLE_NAME).Select("*").Where("id = ?", id).First(&user)
	err = query.Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func FindUsersByName(name string, limit int64, offset int64) (users []User, err error) {
	query := config.DB.Table(TABLE_NAME).Select("*").Where("name = ?", name).Find(&users).Limit(limit).Offset(offset)
	err = query.Error
	if err != nil {
		return users, err
	}
	return users, err
}

func FindUsersByNameLike(name string, limit int64, offset int64) (users []User, err error) {
	query := config.DB.Table(TABLE_NAME).Select("*").Where("name like ?", name).Find(&users).Limit(limit).Offset(offset)
	err = query.Error
	if err != nil {
		return users, err
	}
	return users, err
}
