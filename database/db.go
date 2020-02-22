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
	db.AutoMigrate(&Role{}, &User{}, &Code{}, &Language{})
	db.Model(&Code{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")

	db.Where(&Role{
		Role: "USER",
	}).FirstOrCreate(&Role{})

	db.Where(&Role{
		Role: "ADMIN",
	}).FirstOrCreate(&Role{})

	languages := [10]string{
		"Html",
		"CSS",
		"SCSS",
		"LESS",
		"Javascript",
		"Typescript",
		"Go",
		"C#",
		"Razor",
		"SQL",
	}

	for _, language := range languages {
		db.Where(&Language{
			Language: language,
		}).FirstOrCreate(&Language{})
	}
}

func ClearTables(db *gorm.DB) {
	db.DropTableIfExists(&Code{}, &User{}, &Role{}, &Language{})
}
