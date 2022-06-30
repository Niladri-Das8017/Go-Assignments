package database

import (
	"crudGORM/model"
	"fmt"
)

func FindChats() {

	// var users []model.User
	// DB.Find(&users)
	// for index, user := range users {

	// 	fmt.Printf("%d ID: %v Email: %s	UserName: %s\n", index, user.ID, user.Email, user.UserName)

	// }

	var messages []model.Message
	//db.Where(&model.Message{User: users[0]}).Find(&messages)
	DB.Joins("User").Find(&messages)
	for index, m := range messages {

		fmt.Printf("%d Content: %v UserName: %v\n", index, m.Content, m.User.UserName)
	}

}
