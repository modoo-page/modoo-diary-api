package database

func DeleteKakaoAuth(kakaoId string) (err error) {
	tx := DB.Where("kakao_id = ?", kakaoId).Delete(KakaoAuth{})
	err = tx.Error
	return
}
