package database

import "github.com/jinzhu/gorm"

func CreateLanguage(db *gorm.DB, language *Language) (*Language, error) {
	dbs := db.Create(&language)
	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return language, nil
}

func GetLanguage(db *gorm.DB, language *Language) (*Language, error) {
	dbs := db.First(&language)
	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return language, nil
}

func GetLanguages(db *gorm.DB) (*[]Language, error) {
	var languages []Language
	dbs := db.Find(&languages)
	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &languages, nil
}

func UpdateLanguage(db *gorm.DB, language *Language) (*Language, error) {
	dbs := db.Model(&Language{}).Update(language)
	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return language, nil
}

func DeleteLanguage(db *gorm.DB, language *Language) (*Language, error) {
	dbs := db.Unscoped().Delete(&language)
	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return language, nil
}
