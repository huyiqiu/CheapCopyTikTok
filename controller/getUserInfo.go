package controller

import (
	"minitiktok/service"
	"minitiktok/utils"
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
	_, err2 := utils.ValidateToken(userToken)
	if err2 != nil {
		c.JSON(http.StatusOK, &ErrorResponce{
			StatusCode: 1,
			StatusMsg:  "token过期,请重新登录",
		})
	}
	c.JSON(http.StatusOK, QueryUserInfo(userId, userToken))
}
