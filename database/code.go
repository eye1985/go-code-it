package database

import (
	"github.com/jinzhu/gorm"
)

func QueryUserCode(db *gorm.DB, userId int, codeId int) (*CodeAndLanguage, error) {
	var code CodeAndLanguage

	dbs := db.
		Table("codes").
		Select("*").
		Joins("join languages on languages.id = codes.language_id").
		Where("codes.id = ? AND codes.user_id = ?", codeId, userId).
		First(&code)

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
	dbs := db.
		Table("codes").
		Where("id = ? AND user_id = ?", codeId, userId).
		First(&deleteCode).
		Unscoped().
		Delete(&deleteCode)

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &deleteCode, nil
}
