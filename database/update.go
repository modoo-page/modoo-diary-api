package database

import "gopkg.in/guregu/null.v4"

func UpdateToken(email string, token null.String, time null.Time) (err error) {
	tx := DB.Model(User{}).Where("email = ?", email).Update("auth_token", token).Update("auth_expried_at", time)
	err = tx.Error
	return
}

func UpdateNickname(userId int, nickname string) (err error) {
	tx := DB.Model(User{}).Where("user_id = ?", userId).Update("nickname", nickname)
	err = tx.Error
	return
}
