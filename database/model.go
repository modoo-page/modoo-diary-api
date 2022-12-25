package database

import "time"

type KakaoAuth struct {
	KakaoAuthId int    `json:"kakao_auth_id"`
	KakaoId     string `json:"kakao_id"`
	UserId      int    `json:"user_id"`
}
type Diary struct {
	DiaryId      int       `json:"diary_id"`
	UserId       int       `json:"user_id"`
	DiaryContent string    `json:"diary_content"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
type Hashtag struct {
	HashtagId   int       `json:"hashtag_id"`
	HashtagName string    `json:"hashtag_name"`
	CreatedAt   time.Time `json:"created_at"`
}
type HashtagItem struct {
	HashtagItemId int       `json:"hashtag_item_id"`
	DiaryId       int       `json:"diary_id"`
	HashtagId     int       `json:"hashtag_id"`
	CreatedAt     time.Time `json:"created_at"`
}
type User struct {
	UserId        int       `json:"user_id"`
	Email         string    `json:"email"`
	CreatedAt     time.Time `json:"created_at"`
	IsDeleted     int       `json:"is_deleted"`
	AuthToken     string    `json:"auth_token"`
	AuthExpiredAt time.Time `json:"auth_expired_at"`
}
