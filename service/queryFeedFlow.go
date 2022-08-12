package service

import (
	//"fmt"
	"fmt"
	"minitiktok/repository"
	"minitiktok/utils"
	"minitiktok/utils/logger"
)


func QueryFeedFlow(lastTime string, userToken string) ([]*repository.Video, error){
	videoDao := repository.NewVideoDaoInstance()
	videoList, err := videoDao.QueryVideoList()
	if err != nil {
		println("feed flow goes run")
		return nil, err
	}
	logger.Info(fmt.Sprintf("query feed from %s", userToken))
	// TO DO 根据token验证用户信息
	if userToken != "null" {
		VideoRelationship(userToken, videoList)
	}
	
	return videoList, nil
}


// 查询发布列表
func QueryPublishFlow(userId int, userToken string) ([]*repository.Video, error) {
	videoDao := repository.NewVideoDaoInstance()
	videoList, err := videoDao.QueryPublishList(userId)
	if err != nil {
		println("publish flow goes run")
		return nil, err
	}
	if userToken != "null" {
		VideoRelationship(userToken, videoList)
	}
	return videoList, err
}

// 判断视频的作者我是否关注，判断视频我是否点赞
func VideoRelationship(userToken string, videoList []*repository.Video) {
	for v := range(videoList) {
		// 判断该视频是否被登录用户点赞
		userId, _ := utils.ValidateToken(userToken)
		videoId := videoList[v].Id
		videoList[v].IsFavorite = repository.IsFavorite(userId, videoId)
		authorId := videoList[v].UserId
		isfollow := repository.IsFollow(userId, authorId)
		videoList[v].User.IsFollow = isfollow
	}
}