package controller

import (
	"minitiktok/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"status_msg"`
	UserInfo   interface{} `json:"user"`
}

func QueryUserInfo(userId int, userToken string) *UserInfoResponse {
	user := service.GetUserInfo(userId, userToken)
	if user == nil {
		return &UserInfoResponse{
			StatusCode: 1,
			Msg:        "failed",
		}
	}
	return &UserInfoResponse{
		StatusCode: 0,
		Msg:        "success",
		UserInfo:   user,
	}
}

func UserInfo(c *gin.Context) {
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		println("something goes wrong")
	}
	userToken := c.Query("token")
	c.JSON(http.StatusOK, QueryUserInfo(userId, userToken))
}
