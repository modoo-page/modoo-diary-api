package database

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type KakaoAuth struct {
	KakaoAuthId int `gorm:"primaryKey"`
	KakaoId     string
	UserId      int
	CreatedAt   time.Time
}
type Diary struct {
	DiaryId      int `gorm:"primaryKey"`
	UserId       int
	DiaryContent string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type Hashtag struct {
	HashtagId   int `gorm:"primaryKey"`
	HashtagName string
	CreatedAt   time.Time
}
type HashtagItem struct {
	HashtagItemId int `gorm:"primaryKey"`
	DiaryId       int
	HashtagId     int
	CreatedAt     time.Time
}
type User struct {
	UserId        int `gorm:"primaryKey"`
	Email         string
	Nickname      string
	UserToken     string
	AuthToken     null.String
	AuthExpiredAt null.Time
	IsDeleted     int
	CreatedAt     time.Time
}
