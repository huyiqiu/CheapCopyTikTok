package repository

import (
	"sync"
	"time"
)

const (
	MMDD = "01-02"
)

type Comment struct {
	Id         int       `gorm:"column:id" json:"id"`
	VideoId    int       `gorm:"column:video_id" json:"-"`
	UserId     int       `gorm:"user_id"`
	User       *User     `json:"user"`
	Content    string    `gorm:"content" json:"content"`
	CreateDate time.Time `gorm:"create_date" json:"create_date"`
}

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDaoInstance() *CommentDao {
	commentOnce.Do(func() {
		db.AutoMigrate(&Comment{})
		commentDao = &CommentDao{}
	})
	return commentDao
}

func (*CommentDao) CreateComment(userId int, videoId int, content string) int {
	comment := &Comment{
		VideoId:    videoId,
		UserId:     userId,
		Content:    content,
		CreateDate: time.Now(),
	}

	db.Create(&comment)
	commentId := comment.Id
	return commentId
}

func (*CommentDao) QueryTheComment(commentId int) (*Comment, error) {
	var comment *Comment
	err := db.Preload("User").First(&comment, commentId).Error
	if err != nil {
		println("query failed..")
		return nil, err
	}
	return comment, nil
}

func (*CommentDao) QueryCommentList(videoId int) ([]*Comment, error) {
	var comments []*Comment
	err := db.Preload("User").Where("video_id = ?", videoId).Order("comments.create_date desc").Find(&comments).Error
	if err != nil {
		println("queryCommentList failed")
		return nil, err
	}
	for v := range comments {
		comments[v].CreateDate.Format(MMDD)
	}
	return comments, nil
}

func (*CommentDao) DeleteComment(commentId int) (*Comment, error) {
	err := db.Delete(&Comment{}, commentId).Error
	if err != nil {
		println("delete comment failed")
		return nil, err
	}
	return nil, nil
}
