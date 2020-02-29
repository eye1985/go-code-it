package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateUser(db *gorm.DB, user *User) (*User, error) {
	var role Role
	dbr := db.Where(&Role{Role: "USER"}).First(&role)
	if dbr.Error != nil {
		return nil, dbr.Error
	}

	user.RoleID = &role.ID

	dbu := db.Create(user)
	if dbu.Error != nil {
		return nil, dbu.Error
	}

	return user, nil
}

func Query(db *gorm.DB, q interface{}) error {
	if result := db.Find(q); result.Error != nil {
		return result.Error
	}

	return nil
}

//func QueryOne(db *gorm.DB, q string, qs string, i interface{}) error {
//	if result := db.
//		Where(q, qs).
//		First(&i); result.Error != nil {
//		return result.Error
//	}
//
//	return nil
//}

type Test struct {
	Username string
	Role     string
}

func GetUserAndRole(db *gorm.DB, user *User) (*UserAndRole, error) {
	var foundUser UserAndRole

	dbRes := db.Table("users").
		Where(&user).
		Select("users.*,roles.role").
		Joins("join roles on roles.id = users.role_id").
		Scan(&foundUser)

	if dbRes.Error != nil {
		return nil, dbRes.Error
	}

	return &foundUser, nil
}

func GetUser(db *gorm.DB, user *User) (*User, error) {
	var foundUser User

	dbRes := db.
		Table("users").
		Where(&user).
		Scan(&foundUser)

	if dbRes.Error != nil {
		return nil, dbRes.Error
	}

	return &foundUser, nil
}

func UpdateUser(db *gorm.DB, userId string, u *User) (*User, error) {
	var user User
	dbs := db.Table("users").
		Where("id = ?", userId).
		Find(&user).
		Updates(u)

	if dbs.Error != nil {
		return nil, dbs.Error
	}

	return &user, nil
}

func DeleteUser(db *gorm.DB, user *User) error {
	if r := db.Unscoped().Delete(user); r.Error != nil {
		return r.Error
	}

	return nil
}
