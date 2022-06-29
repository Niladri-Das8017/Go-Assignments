package model

import "gorm.io/gorm"

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

//here, on the Message object, there is both a UserID and ChannelID as well as User and Channel. By default, the
//UserID and ChannelID implicitly used to create a foreign key relationship between the Messege, User and Channel tables, and
//thus must be included in the Messege struct in order to fill the User and the Channel inner struct.
type Message struct {
	gorm.Model
	Content   string
	UserID    uint
	UserName  string
	ChannelID uint
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Channel   Channel `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
