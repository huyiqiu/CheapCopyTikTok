package service

import (
	"bytes"
	"context"
	"fmt"
	"minitiktok/repository"
	"minitiktok/utils"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)


func DoPublish(data []byte, token string, title string) int {
	url := Upload(data)
	userId := utils.ValidateToken(token)
	videoDao := repository.NewVideoDaoInstance()
	// TO DO 截取封面
	coverUrl := "http://rcmz8xyya.hd-bkt.clouddn.com/test/cover.png"
	
	videoDao.CreateNewVideo(userId, url, coverUrl, title)
	return 0
}

func Upload(data []byte) string {
	domain := "http://rcmz8xyya.hd-bkt.clouddn.com"
	accessKey := "WGjh0GoD6PCZNmQt5QD80cbkaK77NK1R7LryoD52"
	secretKey := "i-hBzqxJnOfzGRN0R7DqD5IxIfgPcdexr6apghZe"
	bucket := "2147483648"

	// 使用 returnBody 自定义回复格式
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	dataLen := int64(len(data))
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	url := domain + ret.Key
	fmt.Println(ret.Key, ret.Hash)
	fmt.Println(url)
	return url
}