package service

import (
	"fmt"
	"minitiktok/repository"
	// "minitiktok/utils"
)



func GetUserInfo(userId int, userToken string) *repository.User {
	userDao := repository.NewUserDaoInstance()
	user, err := userDao.QueryUser(userId)
	fmt.Printf("在查看用户信息时调用UserDAO->%p\n", userDao)
	if err != nil {
		println("query user failed")
		return nil
	}
	// 登录用户的ID
	// myId := utils.ValidateToken(userToken)
	// 登录用户与userId的关系
	// user.IsFollow = IsFollow(myId, userId)
	return user
}
