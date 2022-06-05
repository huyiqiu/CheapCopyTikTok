package service

import (
	"minitiktok/repository"
	// "minitiktok/utils"
)

func GetUserInfo(userId int, userToken string) (*repository.User) {
	userDao := repository.NewUserDaoInstance()
	user, err := userDao.QueryUserInfo(userId)
	if err != nil {
		println("query user failed")
		return nil
	}
	// myId := utils.ValidateToken(userToken)
	// user.IsFollow = IsFollow(myId, userId)
	return user
}