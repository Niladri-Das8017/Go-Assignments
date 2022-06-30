package database

import (
	"crudGORM/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func IntiDB() (*gorm.DB, error) {

	//database
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//AutoMigrate
	db.AutoMigrate(&model.Channel{}, &model.User{}, &model.Message{})
	//seeding some initial data into database
	seed(db)

	return db, nil
}

//Insert data to DB
func seed(db *gorm.DB) {

	channels := []model.Channel{
		{Name: "General", Description: "General Discussions"},
		{Name: "Off-Topic", Description: "Wired Stuff goes here"},
		{Name: "Suggestions", Description: "Suggestions goes here"},
	}
	//insert data into channel table
	for _, ch := range channels {
		db.Create(&ch)
	}

	users := []model.User{
		{Email: "niladri@dev.com", UserName: "niladri404"},
		{Email: "bumba@dev.com", UserName: "bumba007"},
	}
	//insert data into channel table
	for _, u := range users {
		db.Create(&u)
		//fmt.Println("ID = ", u.ID)
	}

	var genChat, suggestionsChat model.Channel
	// SELECT * FROM users WHERE Name = 'Gneral';
	db.First(&genChat, "Name = ?", "General") //Inline query
	db.First(&suggestionsChat, "Name = ?", "Suggestions")

	var niladri, bumba model.User
	// SELECT * FROM users WHERE UserName = 'Niladri Das';
	//db.First(&niladri, "UserName = ?", "niladri404")
	db.Where(&model.User{UserName: "niladri404"}).First(&niladri)
	db.Where(&model.User{UserName: "bumba007"}).First(&bumba)

	messages := []model.Message{
		{Content: "Hello", Channel: genChat, User: niladri},
		{Content: "What's up!", Channel: genChat, User: bumba},
		{Content: "You need to push your content to github daily", Channel: suggestionsChat, User: bumba},
	}
	//insert data into channel table
	for _, m := range messages {
		db.Create(&m)
	}

}
