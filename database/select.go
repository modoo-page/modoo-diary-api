package database

func SelectUserByEmail(email string) (user User, err error) {
	tx := DB.Table("user").Where("email=?", email).Scan(&user)
	err = tx.Error
	return
}

type DiaryResponse struct {
	Diary
	User
}

func SelectDiaryListTop10() (diaryList []DiaryResponse, err error) {
	tx := DB.Table("diary").Joins("JOIN user ON user.user_id = diary.user_id").Order("diary_id DESC").Limit(10).Scan(&diaryList)
	err = tx.Error
	return
}
func SelectDiaryListByPaging(page int, length int) (diaryList []DiaryResponse, err error) {
	tx := DB.Table("diary").Joins("JOIN user ON user.user_id = diary.user_id").Order("diary_id DESC").Offset((page - 1) * length).Limit(length).Scan(&diaryList)
	err = tx.Error
	return
}
func SelectDiaryListTop10ByUserId(userId int) (diaryList []Diary, err error) {
	tx := DB.Table("diary").Where("user_id = ?", userId).Order("diary_id DESC").Limit(10).Scan(&diaryList)
	err = tx.Error
	return
}

func SelectUserByKakaoId(kakaoId string) (user User, err error) {
	tx := DB.Select("user_id").Table("kakao_auth").Where("kakao_id", kakaoId).Take(&user)
	err = tx.Error
	return
}

func SelectLoginToken(email string, token string) (user User, err error) {
	tx := DB.Table("user").Where("email = ?", email).Where("auth_token = ?", token).Take(&user)
	err = tx.Error
	return
}

func SelectUserToken(userToken string) (user User, err error) {
	tx := DB.Table("user").Where("user_token = ?", userToken).Take(&user)
	err = tx.Error
	return
}
