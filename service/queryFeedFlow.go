package service

import (
	//"fmt"
	"minitiktok/repository"
)


func QueryFeedFlow(lastTime string, userToken string) ([]*repository.Video, error){
	videoDao := repository.NewVideoDaoInstance()
	videoList, err := videoDao.QueryVideoList()
	if err != nil {
		println("feed flow goes run")
		return nil, err
	}
	// // TO DO 根据token验证用户信息
	// if userToken != "null" {
	// 	for v := range(videoList) {
	// 		//TO DO 判断该视频是否被登录用户点赞，是否关注了视频作者
	
	// 	}
	// }
	
	return videoList, nil
}
