package main

import (
	"crudGORM/database"
	"crudGORM/model"
	"fmt"
	"log"
)

func main() {

	//Initialize Database
	db, err := database.IntiDB()
	if err != nil {
		log.Fatal(err)
	}

	var users []model.User
	db.Find(&users)
	for index, user := range users {

		fmt.Printf("%d ID: %v Email: %s	UserName: %s\n", index, user.ID, user.Email, user.UserName)

	}

	var messages []model.Message
	//db.Where(&model.Message{User: users[0]}).Find(&messages)
	db.Joins("User").Find(&messages)
	for index, m := range messages {

		fmt.Printf("%d Content: %v UserID: %v\n", index, m.Content, m.UserID)
	}

}
