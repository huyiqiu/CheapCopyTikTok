package repository

import (
	"sync"

	"gorm.io/gorm"
)

type Follow struct {
	UserId     int `gorm:"column:user_id"`
	FollowerId int `gorm:"column:follower_id"`
}

func IsFollow(userId int, followerId int) bool {
	var follows []*Follow
	db.Where("user_id = ? and follower_id = ?", userId, followerId).Find(&follows)
	return len(follows) != 0
}

type FollowDao struct {
}

var followDao *FollowDao
var followOnce sync.Once

func NewFollowDaoInstance() *FollowDao {
	followOnce.Do(func(){
		db.AutoMigrate(&Follow{})
		followDao = &FollowDao{}
	})
	return followDao
}

func (*FollowDao) CreateFollow(followerId int, userId int) error {
	follow := &Follow{
		UserId: userId,
		FollowerId: followerId,
	}
	db.Create(&follow)
	user := &User{
		Id: userId,
	}
	db.Model(&user).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1))
	follewer := &User{
		Id: followerId,
	}
	db.Model(&follewer).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1))
	return nil
}

func (*FollowDao) CancelFollow(followerId int, userId int) error {
	follow := &Follow{
		UserId: userId,
		FollowerId: followerId,
	}
	db.Where("user_id = ? and follower_id = ?", userId, followerId).Delete(&follow)
	user := &User{
		Id: userId,
	}
	db.Model(&user).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1))
	follewer := &User{
		Id: followerId,
	}
	db.Model(&follewer).UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1))
	return nil
}

func (*FollowDao) QueryFollowList(userId int) ([]*User, error) {
	var users []*User
	db.Raw("select * from users as u left join follows as f on u.id = f.follower_id where u.id = ?", userId).Scan(&users)
	// isFollow
	for v := range(users){
		users[v].IsFollow = true
	}
	return users, nil
}

func (*FollowDao) QueryFansList(userId int) ([]*User, error) {
	var users []*User
	db.Raw("select * from users as u left join follows as f on u.id = f.user_id where u.id = ?", userId).Scan(&users)
	// isFollow
	for v := range(users){
		users[v].IsFollow = IsFollow(userId, users[v].Id)
	}
	return users, nil
}