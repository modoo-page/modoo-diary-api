package database

import "gorm.io/gorm/clause"

func InsertUser(email string) (err error) {
	var user User
	user.Email = email
	tx := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user) // IF NOT EXIST
	err = tx.Error
	return
}

func InsertKakaoAuth(kakaoId string, userId int) (err error) {
	var kakaoAuth KakaoAuth
	kakaoAuth.KakaoId = kakaoId
	kakaoAuth.UserId = userId
	tx := DB.Create(&kakaoAuth)
	err = tx.Error
	return
}

func InsertDiary(userId int, text string) (err error) {
	var diary Diary
	diary.UserId = userId
	diary.DiaryContent = text
	tx := DB.Create(&diary)
	err = tx.Error
	return
}
