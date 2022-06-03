package repository

import (
	"sync"
	"time"
)

type Video struct {
	Id            int    `gorm:"column:id" json:"id"`
	UserId        int       `gorm:"column:user_id" json:"-"`
	User          User      `gorm:"foreignKey:UserId" json:"author"`
	PlayUrl       string    `gorm:"column:play_url" json:"play_url"`
	CoverUrl      string    `gorm:"column:cover_url" json:"cover_url"`
	Title         string    `gorm:"column:title" json:"title"`
	FavoriteCount int       `gorm:"column:favorite_count" json:"favorite_count"`
	CommentsCount int       `gorm:"column:comments_count" json:"comment_count"`
	CreateTime    time.Time `gorm:"column:created_time" json:"-"`
	IsFavorite    bool      `json:"is_favorite"`
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
