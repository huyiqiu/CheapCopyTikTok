package main

import (
	"minitiktok/controller"
	"minitiktok/repository"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/douyin/feed", controller.Feed)

	r.GET("/douyin/publish/list/", controller.PublishFlow)

	r.GET("/douyin/user/", controller.UserInfo)

	r.GET("/douyin/comment/list/", controller.CommentList)

	r.GET("/douyin/favorite/list/", controller.FavoriteList)

	r.GET("/douyin/relation/follow/list/", controller.FollowList)

	r.GET("/douyin/relation/follower/list/", controller.FansList)

	r.POST("/douyin/user/login/", controller.Login)

	r.POST("/douyin/user/register/", controller.Register)

	r.POST("/douyin/publish/action/", controller.Publish)

	r.POST("/douyin/comment/action/", controller.DoComment)

	r.POST("/douyin/favorite/action/", controller.Favorite)

	r.POST("/douyin/relation/action/", controller.Follow)

	err := r.Run()
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	return nil
}
