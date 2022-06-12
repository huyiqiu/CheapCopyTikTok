package service

import (
	"minitiktok/repository"
	"minitiktok/utils"
)

func FollowOp(token string, userId int, action int) error {
	followerId := utils.ValidateToken(token)
	followDao := repository.NewFollowDaoInstance()
	if action == 1 {
		err := followDao.CreateFollow(followerId, userId)
		if err != nil {
			return err
		}
		return nil
	}
	err := followDao.CancelFollow(followerId, userId)
	if err != nil {
		return err
	}
	return nil
}

func QueryFollowList(userId int, token string) ([]*repository.User, error){
	followDao := repository.NewFollowDaoInstance()
	users, err := followDao.QueryFollowList(userId)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func QueryFansList(userId int, token string) ([]*repository.User, error) {
	followDao := repository.NewFollowDaoInstance()
	users, err := followDao.QueryFansList(userId)
	if err != nil {
		return nil, err
	}
	return users, nil
}