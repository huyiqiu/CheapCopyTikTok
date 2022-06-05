package repository

import (
	"gorm.io/gorm"
)

type Favorite struct {
	UserId  int `gorm:"column:user_id"`
	VideoId int `gorm:"column:video_id"`
}

func IsFavorite(userId int, videoId int) bool {
	var fav *Favorite
	err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(&fav).Error
	return err != gorm.ErrRecordNotFound
}
