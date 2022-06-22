package model

import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
}