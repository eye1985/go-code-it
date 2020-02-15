package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Connect(host string, port string, dbname string, user string, pass string) *gorm.DB {
	dbInfo := fmt.Sprintf("host=%v port=%v sslmode=disable dbname=%v user=%v password=%v", host, port, dbname, user, pass)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Code{})
}

func ClearTables(db *gorm.DB) {
	db.DropTableIfExists(&User{}, &Code{})
}
