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

	r.POST("/douyin/user/login/", controller.Login)

	r.POST("/douyin/user/register/", controller.Register)

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
