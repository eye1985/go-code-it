package dummyData

import (
	"github.com/jinzhu/gorm"
	"postgres/database"
)

func InsertDummyData(db *gorm.DB) {
	typeOfCode := "js"
	code := "alert(123)"

	database.CreateUser(db, "Stein", "stein@gmail.com", []database.Code{
		{
			Title: "My block of codes",
			Type:  &typeOfCode,
			Code:  &code,
		},
		{
			Title: "Block 2",
		},
	})

	database.CreateUser(db, "Are", "are@gmail.com", []database.Code{
		{
			Title: "My block of codes",
		},
	})

	database.CreateUser(db, "Lise", "lise@gmail.com", []database.Code{
		{
			Title: "console.log(123)",
		},
	})

	database.CreateUser(db, "Mona", "mona@gmail.com", nil)
	database.CreateUser(db, "Per", "per@gmail.com", nil)

	users := []database.User{}
	codes := []database.Code{}

	database.Query(db, &users)
	database.Query(db, &codes)
}
