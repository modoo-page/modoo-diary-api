package database

func SelectUserByEmail(email string) (user User, err error) {
	tx := DB.Table("user").Where("email=?", email).Scan(&user)
	err = tx.Error
	return
}
func SelectDiaryListTop10() (diaryList []Diary, err error) {
	tx := DB.Table("diary").Order("diary_id DESC").Limit(10).Scan(&diaryList)
	err = tx.Error
	return
}

func SelectDiaryListTop10ByUserId(userId int) (diaryList []Diary, err error) {
	tx := DB.Table("diary").Where("user_id = ?", userId).Order("diary_id DESC").Limit(10).Scan(&diaryList)
	err = tx.Error
	return
}

func SelectUserIdByKakaoId(kakaoId string) (userId int, err error) {
	tx := DB.Select("user_id").Table("kakao_auth").Where("kakao_id", kakaoId).Take(&userId)
	err = tx.Error
	return
}

func SelectLoginToken(email string, token string) (userId int, err error) {
	tx := DB.Select("user_id").Table("user").Where("email = ?", email).Where("auth_token = ?", token).Scan(&userId)
	err = tx.Error
	return
}
