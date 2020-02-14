package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateUser(db *gorm.DB, name string, email string, codes []Code) error {
	user := User{
		Name:  name,
		Email: email,
		Codes: codes,
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

func GetAssociated(db *gorm.DB, u *User, c *[]Code) {
	db.Model(u).Related(c)
}

func UpdateUser(db *gorm.DB, name string, u *User, codeName string, typeName string, code string) {
	db.Where(&User{Name: name}).First(u)

	res := append(u.Codes, Code{
		Name: codeName,
		Type: &typeName,
		Code: &code,
	})

	u.Codes = res
	db.Save(u)
}

func DeleteUser(db *gorm.DB, user *User) error {
	if r := db.Delete(user); r.Error != nil {
		return r.Error
	}

	return nil
}
