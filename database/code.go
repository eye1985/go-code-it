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

func QueryUserCodes(db *gorm.DB, userId int) (*User, error) {
	var user User
	if dbs := db.Preload("Codes").Table("users").Where("id = ?", userId).Find(&user); dbs.Error != nil {
		return nil, dbs.Error
	}

	return &user, nil
}

func QueryUserCode(db *gorm.DB, userId int, codeId int) (*Code, error) {
	var code Code

	dbs := db.Table("codes").Where("id = ? AND user_id = ?", codeId, userId)

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &code, nil
}

func CreateUserCode(db *gorm.DB, userId string, code *Code) (*Code, error) {
	dbs := db.Where("id = ?", userId).
		First(&User{}).
		Table("codes").
		Create(&code)
	if dbs.Error != nil {
		return nil, dbs.Error
	}
	return code, nil
}

func UpdateUserCode(db *gorm.DB, codeId int, userId int, code *Code) (*Code, error) {
	var updatedCode Code

	dbs := db.
		Table("codes").
		Where("id = ? AND user_id = ?", codeId, userId).
		First(&updatedCode).
		Updates(&code)

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &updatedCode, nil
}

func DeleteUserCode(db *gorm.DB, codeId int, userId int) (*Code, error) {
	var deleteCode Code
	dbs := db.Table("codes").Where("id = ? AND user_id = ?", codeId, userId).First(&deleteCode).Unscoped().Delete(&deleteCode)

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &deleteCode, nil
}
