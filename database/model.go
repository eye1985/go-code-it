package database

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);"`
	Email    string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Codes    []Code `json:"codes"`
}

type Code struct {
	gorm.Model
	Title  string  `json:"title" gorm:"type:varchar(60)"`
	Type   *string `json:"type"`
	Code   *string `json:"code"`
	UserID uint
}

type UserAndCode struct {
	Userid   int    `json:"userId"` //Must be lowercase id for alis to work
	Codeid   int    `json:"codeId"` //Must be lowercase id for alis to work
	Title    string `json:"title"`
	Username string `json:"username"`
}
