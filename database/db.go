package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func Connect(host string, port string, dbname string, user string, pass string) (*gorm.DB, error) {
	dbInfo := fmt.Sprintf("host=%v port=%v sslmode=disable dbname=%v user=%v password=%v", host, port, dbname, user, pass)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Code{}, &Language{})
}

func ClearTables(db *gorm.DB) {
	db.DropTableIfExists(&User{}, &Code{}, &Language{})
}
