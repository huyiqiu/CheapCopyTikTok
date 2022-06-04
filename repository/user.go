package repository

import (
	"gorm.io/gorm"
	"minitiktok/utils"
	"sync"
	"time"
)
// user结构体
type User struct {
	Id            int       `gorm:"column:id" json:"id"`
	Name          string    `gorm:"column:name" json:"name"`
	Pwd string `gorm:"-" json:"-"`
	FollowCount   int       `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int       `gorm:"column:follower_count" json:"follower_count"`
	CreateTime    time.Time `gorm:"column:created_time" json:"-"`
	IsFollow      bool      `json:"is_follow"`
}

/*
 * 使用一个结构体来表示DAO，然后实现DAO的方法
**/
type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			db.AutoMigrate(&User{})
			userDao = &UserDao{}
		})
	return userDao
}

// 创建一个新用户
func (*UserDao)CreateUser(name string, pwd string) {
	user := User{
		Name: name,
		Pwd : pwd,
		CreateTime: time.Now(),
	}
	db.Create(&user)
}

// 实现通过用户名查询用户
func (*UserDao) QueryUserByName(name string) (*User, error) {
	var user *User
	err := db.Where("name = ?", name).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		utils.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return user, nil
}
