package service

import (
	"fmt"
	"minitiktok/repository"
	"minitiktok/utils"
	"minitiktok/utils/logger"
)


const(
	UserNotFound = 1
	PasswordWrong = 2
	UserExist = 3
	OtherERROR = 4
)

func Register(username string, pwd string) (code int, id int, token string) {
	userDao := repository.NewUserDaoInstance()
	user, err := userDao.QueryUserByName(username)
	if user != nil {
		return UserExist, -1, ""
	}
	if err != nil {
		return OtherERROR, -1, ""
	}
	code = 0
	id = userDao.CreateUser(username, pwd)
	token = utils.GenerateToken(id)
	return
}

func Login(username string, pwd string) (code int, id int, token string) {
	userDao := repository.NewUserDaoInstance()
	user, err := userDao.QueryUserByName(username)
	fmt.Printf("在登录时调用UserDAO->%p\n", userDao)
	if user == nil {
		return UserNotFound, -1, ""
	}
	if err != nil {
		logger.Error("something goes wrong when query userInfo")
		return OtherERROR, -1, ""
	}
	password := user.Pwd
	if pwd != password {
		return PasswordWrong, -1, ""
	}
	code = 0
	id = user.Id
	token = utils.GenerateToken(user.Id)
	return
}
