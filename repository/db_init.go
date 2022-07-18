package repository

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var client *redis.Client

func Init() error {
	var err error
	//使用gorm加载数据库
<<<<<<< HEAD
	dsn := "root:123456@tcp(127.0.0.1:3306)/minitiktok?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	client = redis.NewClient(&redis.Options{
		Addr: "47.113.185.142:6379", 
=======
	dsn := "root:123456@tcp(172.17.251.27:3306)/minitiktok?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	client = redis.NewClient(&redis.Options{
		Addr: "172.17.251.27:6379", //远程，上线后应当切换到内网ip
>>>>>>> 42ed5dcc955aeaf0e62675cbe2b4d57c17be755f
		Password: "147998",
		DB: 0,
	})
	return err
}
