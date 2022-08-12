package controller

import (
	"minitiktok/service"
	"net/http"
	"strconv"
	"minitiktok/utils"
	"github.com/gin-gonic/gin"
)

type CommentResponce struct {
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	Comment    interface{} `json:"comment"`
}

type CommentListResponce struct {
	StatusCode  int         `json:"status_code"`
	StatusMsg   string      `json:"status_msg"`
	CommentList interface{} `json:"comment"`
}

type ErrorResponce struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func CommentInfo(token string, videoId int, action int, content string, commentId int) *CommentResponce {
	comment, err := service.CommentOpt(token, videoId, action, content, commentId)
	if err != nil {
		return &CommentResponce{
			StatusCode: 1,
			StatusMsg:  "failed",
		}
	}
	return &CommentResponce{
		StatusCode: 0,
		StatusMsg:  "success",
		Comment:    comment,
	}
}

func CommentListInfo(token string, videoId int) *CommentListResponce {
	commentList, err := service.QueryCommentList(token, videoId)
	if err != nil {
		return &CommentListResponce{
			StatusCode: 1,
			StatusMsg:  "failed",
		}
	}
	return &CommentListResponce{
		StatusCode:  0,
		StatusMsg:   "success",
		CommentList: commentList,
	}
}

func DoComment(c *gin.Context) {
	userToken := c.Query("token")
	_, err := utils.ValidateToken(userToken)
	if err != nil {
		c.JSON(http.StatusOK, &ErrorResponce{
			StatusCode: 1,
			StatusMsg:  "token过期,请重新登录",
		})
	}
	videoId, err := strconv.Atoi(c.Query("video_id"))
	action, err2 := strconv.Atoi(c.Query("action_type"))
	content := c.DefaultQuery("comment_text", "")
	commentId, err3 := strconv.Atoi(c.DefaultQuery("comment_id", "0"))
	if err != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusOK, &CommentResponce{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
	}
	c.JSON(http.StatusOK, CommentInfo(userToken, videoId, action, content, commentId))
}

func CommentList(c *gin.Context) {
	userToken := c.Query("token")
	_, err := utils.ValidateToken(userToken)
	if err != nil {
		c.JSON(http.StatusOK, &ErrorResponce{
			StatusCode: 1,
			StatusMsg:  "token过期,请重新登录",
		})
	}
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, &CommentResponce{
			StatusCode: 1,
			StatusMsg:  "failed",
		})
	}
	c.JSON(http.StatusOK, CommentListInfo(userToken, videoId))
}
