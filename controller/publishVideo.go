package controller

import (
	"io/ioutil"
	"minitiktok/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublishResponce struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"status_msg"`
}

func PublishVideo(data []byte, token string, title string) *PublishResponce {
	code := service.DoPublish(data, token, title)
	if code != 0 {
		return &PublishResponce{
			StatusCode: 1,
			Msg: "failed",
		}
	}
	return &PublishResponce{
		StatusCode: 0,
		Msg: "success",
	}
}

func Publish(c *gin.Context) {
	userToken := c.PostForm("token")
	title := c.PostForm("title")
	file, _ := c.FormFile("data")
	data, _ := file.Open()
	dataInfo, _ := ioutil.ReadAll(data)
	c.JSON(http.StatusOK, PublishVideo(dataInfo, userToken, title))
}