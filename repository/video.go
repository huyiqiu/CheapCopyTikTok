package repository

import (
	"sync"
	"time"
)

type Video struct {
	Id            int       `gorm:"column:id" json:"id"`
	UserId        int       `gorm:"column:user_id" json:"-"`
	User          User      `gorm:"foreignKey:UserId" json:"author"`
	PlayUrl       string    `gorm:"column:play_url" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url" json:"cover_url"`
	Title         string    `gorm:"column:title" json:"title"`
	FavoriteCount int       `gorm:"column:favorite_count" json:"favorite_count"`
	CommentsCount int       `gorm:"column:comments_count" json:"comment_count"`
	CreateTime    time.Time `gorm:"column:created_time" json:"-"`
	IsFavorite    bool      `gorm:"-" json:"is_favorite"`
}

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) QueryVideoList() ([]*Video, error) {
	var videos []*Video
	println("query video")
	db.Limit(30).Order("videos.created_time desc").Preload("User").Find(&videos)
	println(videos)
	return videos, nil
}

// 发布列表
func (*VideoDao) QueryPublishList(userId int) ([]*Video, error) {
	var videos []*Video
	println("query publish list")
	//
	db.Where("user_id = ?", userId).Order("videos.created_time desc").Preload("User").Find(&videos)
	return videos, nil
}

// 创建一个视频记录
func (*VideoDao) CreateNewVideo(userId int, playUrl string, coverUrl string, title string) {
	video := Video{
		UserId:     userId,
		PlayUrl:    playUrl,
		CoverUrl:   coverUrl,
		Title:      title,
		CreateTime: time.Now(),
	}
	err := db.Create(&video).Error
	if err != nil {
		println("create a new video failed")
	}
}
