package database

func SelectUserByEmail(email string) (user User, err error) {
	tx := DB.Table("users").Where("email=?", email).Scan(&user)
	err = tx.Error
	return
}
func SelectDiaryListTop10() (diaryList []Diary, err error) {
	tx := DB.Table("diaries").Limit(10).Scan(&diaryList)
	err = tx.Error
	return
}

func SelectDiaryListTop10ByUserId(userId int) (diaryList []Diary, err error) {
	tx := DB.Table("diaries").Where("userId = ?", userId).Limit(10).Scan(&diaryList)
	err = tx.Error
	return
}

func SelectUserIdByKakaoId(kakaoId string) (userId int, err error) {
	tx := DB.Select("user_id").Table("auth_kakao").Where("kakao_id", kakaoId).Scan(&userId)
	err = tx.Error
	return
}

func SelectLoginToken(token string) (userId int, err error) {
	tx := DB.Select("user_id").Table("user").Where("auth_token", token).Scan(&userId)
	err = tx.Error
	return
}
