package models

import (
	"github.com/kataras/golog"
	"../util"
)

type User struct {
	UserId int64  `json:"user_id"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
}

func (User) TableName() string {
	return "h_user"
}

func (u *User) GetUserInfo(userId string) interface{} {
	var user User
	db := util.GetInstance().GetMysqlDB()
	err := db.Where("user_id = ?", userId).First(&user).Error

	if err != nil {
		golog.Println("查询失败")
	}
	return user
}

func (u *User) GetUserList(limit, offset int) interface{} {
	var users [] User
	db := util.GetInstance().GetMysqlDB()
	db.Offset(offset).Limit(limit).Find(&users)
	return users
}

func (u *User) CreateUser(user *User) {
	db := util.GetInstance().GetMysqlDB()
	db.Table("h_user").Create(&user)
}
