package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"postgres/database"
	"postgres/server"
)

func cleanup(db *gorm.DB) {
	if r := recover(); r != nil {
		fmt.Println("Recovered from: ", r)
		return
	}

	db.Close()
}

func main() {
	db := database.Connect()
	server.Db = db
	defer cleanup(db)

	database.ClearTables(db)
	database.Migrate(db)

	// DB test
	typeOfCode := "js"
	code := "alert(123)"

	database.CreateUser(db, "Arne", "arne@gmail.com", []database.Code{
		{
			Name: "My block of codes",
			Type: &typeOfCode,
			Code: &code,
		},
		{
			Name: "Block 2",
		},
	})

	database.CreateUser(db, "Bjarne", "bjarne@gmail.com", []database.Code{
		{
			Name: "My block of codes",
		},
	})

	users := []database.User{}
	codes := []database.Code{}

	database.Query(db, &users)
	database.Query(db, &codes)

	database.UpdateUser(db, "Bjarne", &users[1], "More block of codes", "hmm", "alert(asdasdasd)")
	database.UpdateUser(db, "Bjarne", &users[1], "More block of codes 2", "hmm", "alert(asdasdasd)")
	database.UpdateUser(db, "Bjarne", &users[1], "More block of codes 3", "hmm", "alert(asdasdasd)")

	for _, user := range users {
		scopedCodes := []database.Code{}
		database.GetAssociated(db, &user, &scopedCodes)

		fmt.Printf("%v \n", user.Name)

		for _, code := range scopedCodes {
			if code.Code != nil {
				fmt.Println(*code.Code)
			}
		}
	}
	// db test end

	server.StartServer()
}
