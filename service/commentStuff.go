package service

import (
	"minitiktok/repository"
	"minitiktok/utils"
)

func CommentOpt(token string, videoId int, action int, content string, commentId int) (*repository.Comment, error) {
	commentDao := repository.NewCommentDaoInstance()
	if action == 1 {
		userId := utils.ValidateToken(token)
		id := commentDao.CreateComment(userId, videoId, content)
		return commentDao.QueryTheComment(id)
	}
	return commentDao.DeleteComment(commentId, videoId) // action == 2
}

func QueryCommentList(token string, videoId int) ([]*repository.Comment, error) {
	// TODO: 判断isFollow
	// userId := utils.ValidateToken(token)
	// 
	commentDao := repository.NewCommentDaoInstance()
	return commentDao.QueryCommentList(videoId)
}