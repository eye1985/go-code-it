package database

import (
	"github.com/jinzhu/gorm"
)

func SearchCodes(db *gorm.DB, query string, offset int16, limit int16) (*uint16, []UserAndCode, error) {
	var res []UserAndCode
	var count uint16

	if cDb := db.
		Table("codes").
		Where("public = ?", true).
		Count(&count); cDb.Error != nil {
		return nil, nil, cDb.Error
	}

	if cDb := db.
		Offset(offset-1).
		Limit(limit).
		Table("users").
		Select("users.id AS userid, codes.id AS codeid, codes.title, codes.description, users.username").
		Joins("join codes ON codes.user_id = users.id").
		Where("public = ?", true).
		Scan(&res); cDb.Error != nil {
		return nil, nil, cDb.Error
	}

	return &count, res, nil
}
