package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Todo struct
type Todo struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	Title     string             `bson:"title" json:"title"`
	Desc      string             `bson:"desc" json:"desc"`
	Completed bool               `bson:"completed" json:"completed"`
}
