package database

import "github.com/jinzhu/gorm"

const dbInfo = "host=127.0.0.1 port=5432 sslmode=disable dbname=codebase user=postgres password=admin"
const dialect = "postgres"

func Connect() *gorm.DB {
	db, err := gorm.Open(dialect, dbInfo)
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
