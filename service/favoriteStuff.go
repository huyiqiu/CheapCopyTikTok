package service

import (
	"minitiktok/repository"
	"minitiktok/utils"
)

func FavoriteOperation(token string, videoId int, action int) error {
	favDao := repository.NewFavDaoInstance()
	userId, _ := utils.ValidateToken(token)
	if action == 1 {
		err := favDao.CreateLike(userId, videoId)
		if err != nil {
			return err
		}
		return nil
	}
	err := favDao.CancelLike(userId, videoId)
	if err != nil {
		return err
	}
	return nil
}

func QueryFavoriteList(userId int, token string) ([]*repository.Video, error) {
	favDao := repository.NewFavDaoInstance()
	favList, err := favDao.QueryFavList(userId)
	if err != nil {
		return nil, err
	}
	return favList, nil
}