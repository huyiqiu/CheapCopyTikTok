package service

import (
	"minitiktok/repository"
	"minitiktok/utils"
)

const (
	NoSuchUser = 1
	OtherError = 2
)

var userDao = repository.NewUserDaoInstance()

// func Register(username string, pwd string) (code int, id int, token string) {

// }

func Login(username string, pwd string) (code int, id int, token string) {
	user, err := userDao.QueryUserByName(username)
	if user == nil {
		return NoSuchUser, -1, ""
	}
	if err != nil {
		utils.Logger.Error("something goes wrong when query userInfo")
		return OtherError, -1, ""
	}
	code = 0
	id = user.Id
	token = utils.GenerateToken(user.Id)
	return
}
