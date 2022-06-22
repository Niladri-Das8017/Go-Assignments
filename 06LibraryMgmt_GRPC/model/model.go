package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//Book Structure
type BookDetails struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Author string             `json:"author,omitempty"`
}
