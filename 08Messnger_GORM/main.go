package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Channel struct {
	gorm.Model
	Name        string
	Description string
}

type User struct {
	gorm.Model
	Email    string
	UserName string
}

type Message struct {
	gorm.Model
	Content   string
	UserID    uint
	UserName string
	ChannelID uint
	User      User
	Channel   Channel
}

func setup(db *gorm.DB) {

	db.AutoMigrate(&Channel{}, &User{}, &Message{})
	seed(db)
}

//Insert data to DB
func seed(db *gorm.DB) {

	channels := []Channel{
		{Name: "General", Description: "General Discussions"},
		{Name: "Off-Topic", Description: "Wired Stuff goes here"},
		{Name: "Suggestions", Description: "Suggestions goes here"},
	}
	//insert data into channel table
	for _, ch := range channels {
		db.Create(&ch)
	}

	users := []User{
		{Email: "niladri@dev.com", UserName: "niladri404"},
		{Email: "bumba@dev.com", UserName: "bumba007"},
	}
	//insert data into channel table
	for _, u := range users {
		db.Create(&u)
	}

	var genChat, suggestionsChat Channel
	// SELECT * FROM users WHERE Name = 'Gneral';
	db.First(&genChat, "Name = ?", "General") //Inline query
	db.First(&suggestionsChat, "Name = ?", "Suggestions")

	var niladri, bumba User
	// SELECT * FROM users WHERE UserName = 'Niladri Das';
	//db.First(&niladri, "UserName = ?", "niladri404")
	db.Where(&User{UserName: "niladri404"}).First(&niladri) 
	db.Where(&User{UserName: "bumba007"}).First(&bumba) 

	messages := []Message{
		{Content: "Hello", Channel: genChat, User: niladri},
		{Content: "What's up!", Channel: genChat, User: bumba},
		{Content: "Yoou need to push your content to giithub daily", Channel: suggestionsChat, User: bumba},
	}
	//insert data into channel table
	for _, m := range messages {
		db.Create(&m)
	}

}

func main() {
	//database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Faild to Connect to the database")
	}

	//calling function
	setup(db)

	var users []User
	db.Find(&users)
	for index, user := range users {

		fmt.Printf("%d ID: %v Email: %s	UserName: %s\n", index, user.ID, user.Email, user.UserName)

	}

	 var messages []Message
	 db.Where(&Message{User: users[0]}).Find(&messages)

	 for index, m := range messages {

		fmt.Printf("%d Content %v UserID: %v\n", index, m.Content, m.UserID)
	 }

 }
