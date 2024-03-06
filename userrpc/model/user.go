package model

import (
	"gorm.io/gorm"
	"zg5/z304/framework/mysql"
)

type User struct {
	gorm.Model
	Tel      string
	Password string
}

func NewUser() *User {
	return new(User)
}

func (u *User) Get(user *User) (info *User, err error) {
	info = new(User)
	err = mysql.DB.Where(user).First(info).Error
	return
}
