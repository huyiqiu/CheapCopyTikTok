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

func Follow(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.Atoi(c.Query("to_user_id"))
	action, _ := strconv.Atoi(c.Query("action_type"))
	c.JSON(http.StatusOK, FollowInfo(userId, token, action))
}