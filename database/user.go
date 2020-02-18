package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateUser(db *gorm.DB, name string, email string, codes []Code) error {
	user := User{
		Username: name,
		Email:    email,
		Codes:    codes,
	}

	if dbc := db.Create(&user); dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

func Query(db *gorm.DB, q interface{}) error {
	if result := db.Find(q); result.Error != nil {
		return result.Error
	}

	return nil
}

func QueryOne(db *gorm.DB, q string, qs string, i interface{}) error {
	if result := db.Where(q, qs).Find(i); result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateUser(db *gorm.DB, userId string, username string, email string) (*User, error) {
	dbs := db.Table("users").Where("id = ?", userId).Updates(User{
		Username: username,
		Email:    email,
	})

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	var user User
	dbs.Scan(&user)

	return &user, nil
}

func SoftDeleteUser(db *gorm.DB, user *User) error {
	if r := db.Delete(user); r.Error != nil {
		return r.Error
	}

	return nil
}
