package database

func UpdateToken(email string, token string) (err error) {
	tx := DB.Model(User{}).Where("email = ?", email).Update("auth_token", token)
	err = tx.Error
	return
}
