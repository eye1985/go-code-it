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

func GetLanguages(db *gorm.DB) (*[]IdAndLanguage, error) {
	var languages []IdAndLanguage

	dbs := db.Find(&[]Language{}).Scan(&languages)

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
	dbs := db.Unscoped().
		Where(&language).
		Delete(&Language{})

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return language, nil
}
