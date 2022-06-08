package controller

import (
	"minitiktok/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentResponce struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Comment    interface{} `json:"comment"`
}

func CommentInfo(token string, videoId int, action int, content string, commentId int) *CommentResponce {
	comment, err := service.CommentOpt(token, videoId, action, content, commentId)
	if err != nil {
		return &CommentResponce{
			StatusCode: 1,
			StatusMsg: "failed",
		}
	}
	return &CommentResponce{
		StatusCode: 0,
		StatusMsg: "success",
		Comment: comment,
	}
}

func DoComment(c *gin.Context) {
	userToken := c.PostForm("token")
	videoId, err := strconv.Atoi(c.PostForm("video_id"))
	action, err2 := strconv.Atoi(c.PostForm("action_type"))
	content := c.DefaultQuery("comment_text", "")
	commentId, err3 := strconv.Atoi(c.DefaultQuery("comment_id", "0"))
	if err != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK ,&CommentResponce{
			StatusCode: 1,
			StatusMsg: "failed",
		})
	}
	c.JSON(http.StatusOK, CommentInfo(userToken, videoId, action, content, commentId))
}