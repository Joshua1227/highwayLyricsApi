package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	Title      string `bson:"title" json:"title"`
	Lyrics     string `bson:"lyrics" json:"lyrics"`
	AddedBy    string `bson:"addedby" json:"added_by"`
	ApprovedBy string `bson:"approvedby" json:"approved_by"`
	// language
	// writer
	// tags
}

type Title struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Title string             `bson:"title" json:"title"`
}
