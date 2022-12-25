package database

import "gorm.io/gorm/clause"

func InsertUser(email string) (err error) {
	var user User
	user.Email = email
	tx := DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user) // IF NOT EXIST
	err = tx.Error
	return
}
