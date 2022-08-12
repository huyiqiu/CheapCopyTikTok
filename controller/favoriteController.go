package controller

import (
	"minitiktok/service"
	"minitiktok/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteOpResponce struct {
	StatusCode int    `json:"status_code"`
	Msg        string `json:"status_msg"`
}

type FavListResponce struct {
	StatusCode int         `json:"status_code"`
	Msg        string      `json:"status_msg"`
	VideoList  interface{} `json:"video_list"`
}


func DoFavorite(token string, videoId int, action int) *FavoriteOpResponce {
	err := service.FavoriteOperation(token, videoId, action)
	if err != nil {
		return &FavoriteOpResponce{
			StatusCode: 1,
			Msg: "failed",
		}
	}
	return &FavoriteOpResponce{
		StatusCode: 0,
		Msg: "success",
	}
}

func FavListInfo(userId int, token string) *FavListResponce {
	favList, err := service.QueryFavoriteList(userId, token)
	if err != nil {
		return &FavListResponce{
			StatusCode: 1,
			Msg: "failed",
		}
	}
	return &FavListResponce{
		StatusCode: 0,
		Msg: "success",
		VideoList: favList,
	}
}

func Favorite(c *gin.Context) {
	token := c.Query("token")
	_, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusOK, &ErrorResponce{
			StatusCode: 1,
			StatusMsg:  "token过期,请重新登录",
		})
	}
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		println("favController goes wrong..")
	}
	action, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		println("favController goes wrong..")
	}
	c.JSON(http.StatusOK, DoFavorite(token, videoId, action))
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	_, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusOK, &ErrorResponce{
			StatusCode: 1,
			StatusMsg:  "token过期,请重新登录",
		})
	}
	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		println("favController goes wrong..")
	}
	c.JSON(http.StatusOK, FavListInfo(userId, token))
}