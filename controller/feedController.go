package controller

import (
	"minitiktok/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)
type FeedResponce struct {
	StatusCode int `json:"status_code"`
	Msg string `json:"status_msg"`
	NextTime int64 `json:"next_time,omitempty"`
	VideoList interface{} `json:"video_list"`
}

func QueryVideoInfo(lastTime string, userToken string) *FeedResponce {
	videoInfo, err := service.QueryFeedFlow(lastTime, userToken)
	if err != nil {
		return &FeedResponce{
			StatusCode: -1,
			Msg: err.Error(),
		}
	}
	return &FeedResponce{
		StatusCode: 0,
		Msg: "success",
		NextTime: time.Now().Unix(),
		VideoList: videoInfo,
	}
}

func Feed(c *gin.Context){
	lastTime := c.DefaultQuery("latest_time", strconv.FormatInt(time.Now().Unix(), 10))
	userToken := c.DefaultQuery("token", "null")
	c.JSON(http.StatusOK, QueryVideoInfo(lastTime, userToken))
}