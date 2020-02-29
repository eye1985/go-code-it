package database

import (
	"codepocket/encrypt"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
)

func InsertDummyData(db *gorm.DB) {
	id := uint(1)

	uname := "abu"
	pass := "password"
	email := "asdasd@asd.com"
	code := fmt.Sprintf(`alert("Hello world")`)
	isPublic := false

	hashedPwd, err := encrypt.HashAndSalt([]byte(pass))
	if err != nil {
		log.Fatal(err)
	}

	db.Create(&User{
		Username: &uname,
		Password: &hashedPwd,
		Email:    &email,
		Codes: []Code{
			{
				Title:       "Private Alert code",
				Description: "yes yes private",
				Code:        &code,
				Public:      &isPublic,
				LanguageID:  &id,
			},

			{
				Title:       "Public Alert code",
				Description: "yes yes public",
				Code:        &code,
				LanguageID:  &id,
			},
		},
		RoleID: &id,
	})

	for i := 0; i < 66; i++ {

		uname := "User" + strconv.Itoa(i)
		pass := "Userpassword" + strconv.Itoa(i)
		email := "email@" + strconv.Itoa(i) + "asdad.com"
		code := fmt.Sprintf(`alert("Hello world %v")`, strconv.Itoa(i))

		hashedPwd, err := encrypt.HashAndSalt([]byte(pass))
		if err != nil {
			log.Fatal(err)
		}

		db.Create(&User{
			Username: &uname,
			Password: &hashedPwd,
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
