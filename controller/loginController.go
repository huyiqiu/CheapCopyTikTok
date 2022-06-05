package controller

import (
	"minitiktok/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginAndRegisterResponce struct {
	StatusCode int `json:"status_code"`
	Msg string `json:"status_msg"`
	UserId int `json:"user_id"`
	Token string `json:"token"`
}

func RegisterInfo(username string, password string) *LoginAndRegisterResponce {
	code, id, token := service.Register(username, password)
	if code == 3 {
		return &LoginAndRegisterResponce{
			StatusCode: 3,
			Msg: "UserAlreadyExist",
		}
	} else if code == 4 {
		return &LoginAndRegisterResponce{
			StatusCode: 4,
			Msg: "something goes wrong..",
		}
	}
	return &LoginAndRegisterResponce{
		StatusCode: 0,
		Msg: "success",
		UserId: id,
		Token: token,
	}
}

func LoginInfo(username string, password string) *LoginAndRegisterResponce {
	code, id, token := service.Login(username, password)
	if code == 1 {
		return &LoginAndRegisterResponce{
			StatusCode: 1,
			Msg: "UserNotFound..",
		}
	} else if code == 2 {
		return &LoginAndRegisterResponce{
			StatusCode: 2,
			Msg: "password wrong..",
		}
	} else if code == 4 {
		return &LoginAndRegisterResponce{
			StatusCode: 4,
			Msg: "something goes wrong..",
		}
	}
	return &LoginAndRegisterResponce{
		StatusCode: 0,
		Msg: "success",
		UserId: id,
		Token: token,
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, LoginInfo(username, password))
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(http.StatusOK, RegisterInfo(username, password))
}