package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Username string  `json:"username" gorm:"type:varchar(100);"`
	Password *string `json:"password" gorm:"type:varchar(30);not null"`
	Email    *string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Codes    []Code  `json:"codes"`
}

type Code struct {
	gorm.Model
	Title      string  `json:"title" gorm:"type:varchar(60)"`
	Type       *string `json:"type"`
	Code       *string `json:"code"`
	UserID     uint
	LanguageID uint
}

type Language struct {
	gorm.Model
	Language string `json:"language" gorm:"type:varchar(60);unique;not null"`
	Codes    []Code
}

func (l *Language) BeforeCreate() (err error) {
	if len(strings.TrimSpace(l.Language)) == 0 {
		err = errors.New("Language cannot be empty")
	}

	return
}

// Used in result
type UserAndCode struct {
	Userid   int    `json:"userId"` //Must be lowercase id for alis to work
	Codeid   int    `json:"codeId"` //Must be lowercase id for alis to work
	Title    string `json:"title"`
	Username string `json:"username"`
}
