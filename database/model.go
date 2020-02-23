package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Username *string `json:"username" gorm:"type:varchar(30);unique;unique_index;not null"`
	Password *string `json:"password" gorm:"type:varchar(100);not null"`
	Email    *string `json:"email" gorm:"type:varchar(100);unique;unique_index;not null"`
	Codes    []Code  `json:"codes"`
	RoleID   *uint   `json:"roleId" gorm:"not null"`
}

type Role struct {
	gorm.Model
	Role  string `json:"role" gorm:"unique;unique_index;not null"`
	Users []User `json:"users"`
}

type Code struct {
	gorm.Model
	Title       string  `json:"title" gorm:"type:varchar(60);not null"`
	Description string  `json:"description" gorm:"type:varchar(300)"`
	Code        *string `json:"code" gorm:"not null"`
	Public      bool    `json:"public" gorm:"default:true"`
	UserID      *uint   `gorm:"not null"`
	LanguageID  *uint   `gorm:"not null"`
}

type Language struct {
	gorm.Model
	Language string `json:"language" gorm:"type:varchar(60);unique;unique_index;not null"`
	Codes    []Code
}

func (l *Language) BeforeCreate() (err error) {
	if len(strings.TrimSpace(l.Language)) == 0 {
		err = errors.New("Language cannot be empty")
	}

	return
}

// Output model

// Used in result
type UserAndCode struct {
	Userid   int    `json:"userId"` //Must be lowercase id for alis to work
	Codeid   int    `json:"codeId"` //Must be lowercase id for alis to work
	Title    string `json:"title"`
	Username string `json:"username"`
}

type UserAndRole struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Pagination struct {
	Codes       []UserAndCode `json:"codes"`
	CurrentPage int16         `json:"currentPage"`
	NextStart   int16         `json:"nextStart"`
	PrevStart   int16         `json:"prevStart"`
	TotalPage   int16         `json:"totalPage"`
}
