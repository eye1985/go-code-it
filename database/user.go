package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateUser(db *gorm.DB, name string, email string, password string) (*User, error) {
	user := User{
		Username: name,
		Password: &password,
		Email:    &email,
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		return nil, dbc.Error
	}

	return &user, nil
}

func Query(db *gorm.DB, q interface{}) error {
	if result := db.Find(q); result.Error != nil {
		return result.Error
	}

	return nil
}

func QueryOne(db *gorm.DB, q string, qs string, i interface{}) error {
	if result := db.
		Where(q, qs).
		Find(i); result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateUser(db *gorm.DB, userId string, username string, email string, password string) (*User, error) {
	var user User
	dbs := db.Table("users").
		Where("id = ?", userId).
		Find(&user).
		Updates(User{
			Username: username,
			Password: &password,
			Email:    &email,
		})

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &user, nil
}

func SoftDeleteUser(db *gorm.DB, user *User) error {
	if r := db.Delete(user); r.Error != nil {
		return r.Error
	}

	return nil
}
