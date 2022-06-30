package database

import (
	"CRUD_GORM/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {

	db, err := gorm.Open(sqlite.Open("C:/Users/Niladri Das/go/Go-Assignments/07CRUD_GORM/database/Phonebook.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Faild to Connect to the database")
	}

	db.AutoMigrate(&model.Contact{}, model.Ph{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Contact{})

	DB = db

}
