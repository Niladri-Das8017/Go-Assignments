package database

import (
	"CRUD_GORM/model"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("PhoneBook.db"), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Faild to Connect to the database")
	}

	db.AutoMigrate(&model.Contact{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Contact{})

	return db, nil
}
