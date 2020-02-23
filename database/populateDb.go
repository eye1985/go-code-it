package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

func InsertDummyData(db *gorm.DB) {

	id := uint(1)

	for i := 0; i < 66; i++ {

		uname := "User" + strconv.Itoa(i)
		pass := "Userpassword" + strconv.Itoa(i)
		email := "email@" + strconv.Itoa(i) + "asdad.com"
		code := fmt.Sprintf(`alert("Hello world %v")`, strconv.Itoa(i))

		db.Create(&User{
			Username: &uname,
			Password: &pass,
			Email:    &email,
			Codes: []Code{
				{
					Title:       "My code sample" + strconv.Itoa(i),
					Description: "Some dexc",
					Code:        &code,
					LanguageID:  &id,
				},
			},
			RoleID: &id,
		})
	}
}
