package database

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(100);"`
	Email string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Codes []Code `json:"codes"`
}

type Code struct {
	gorm.Model
	Name   string  `json:"name" gorm:"type:varchar(60)"`
	Type   *string `json:"type"`
	Code   *string `json:"code"`
	UserID uint
}
