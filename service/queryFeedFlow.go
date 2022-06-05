package service

import (
	//"fmt"
	"minitiktok/repository"
	"minitiktok/utils"
)


func QueryFeedFlow(lastTime string, userToken string) ([]*repository.Video, error){
	videoDao := repository.NewVideoDaoInstance()
	videoList, err := videoDao.QueryVideoList()
	if err != nil {
		println("feed flow goes run")
		return nil, err
	}

	// TO DO 根据token验证用户信息
	if userToken != "null" {
		for v := range(videoList) {
			//TO DO 判断该视频是否被登录用户点赞，是否关注了视频作者
			userId := utils.ValidateToken(userToken)
			videoId := videoList[v].Id
			videoList[v].IsFavorite = repository.IsFavorite(userId, videoId)
		}
	}
	
	return videoList, nil
}


// // 查询发布列表
// func QueryPublishFlow(userId int, userToken string) ([]*repository.Video, error) {
	
// }
