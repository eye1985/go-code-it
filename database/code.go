package database

import (
	"github.com/jinzhu/gorm"
)

func QueryAllCodes(db *gorm.DB) ([]UserAndCode, error) {
	var result []UserAndCode

	if dbs := db.Raw("select c.title, c.id AS codeid ,u.username, u.id AS userid from codes c, users u where c.user_id = u.id"); dbs.Error != nil {
		return nil, dbs.Error
	} else {
		dbs.Scan(&result)
	}

	return result, nil
}

func QueryUserCodes(db *gorm.DB, userId int) (User, error) {
	var user User
	if dbs := db.Preload("Codes").Table("users").Where("id = ?", userId).Find(&user); dbs.Error != nil {
		return user, dbs.Error
	}

	return user, nil
}
