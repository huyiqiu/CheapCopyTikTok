package controller

import (
	"minitiktok/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowResponce struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"status_msg"`
}

type FollowListResponce struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"status_msg"`
	UserList interface{} `json:"user_list"`
}

func FollowInfo(userId int, token string, action int) *FollowResponce {
	err := service.FollowOp(token, userId, action)
	if err != nil {
		return &FollowResponce{
			StatusCode: 1,
			Msg: "failed",
		}
	}
	return &FollowResponce{
		StatusCode: 0,
		Msg: "success",
	}
}

func FollowListInfo(userId int, token string) *FollowListResponce {
	users, err := service.QueryFollowList(userId, token)
	if err != nil {
		return &FollowListResponce{
			StatusCode: 1,
			Msg: "failed",
		}
	}
	return &FollowListResponce{
		StatusCode: 0,
		Msg: "success",
		UserList: users,
	}
}

func FansListInfo(userId int, token string) *FollowListResponce {
	users, err := service.QueryFansList(userId, token)
	if err != nil {
		return &FollowListResponce{
			StatusCode: 1,
			Msg: "failed",
		}
	}
	return &FollowListResponce{
		StatusCode: 0,
		Msg: "success",
		UserList: users,
	}
}

func Follow(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.Atoi(c.Query("to_user_id"))
	action, _ := strconv.Atoi(c.Query("action_type"))
	c.JSON(http.StatusOK, FollowInfo(userId, token, action))
}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.Atoi(c.Query("user_id"))
	c.JSON(http.StatusOK, FollowListInfo(userId, token))
}

func FansList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.Atoi(c.Query("user_id"))
	c.JSON(http.StatusOK, FansListInfo(userId, token))
}