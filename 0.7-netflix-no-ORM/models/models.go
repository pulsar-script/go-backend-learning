package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// custome data type
type Netflix struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie     string             `json:"movie,omitempty" bson:"movie,omitempty"`
	IsWatched bool               `json:"isWatched" bson:"isWatched"`

	//* Recommend way => Keep both naming in JSON & BOSN same

	//? Even if you keep different, json request will work, go and try and read in Notion Notes

	//! IsWatched bool            `json:"isWatched,omitempty" bson:"isWatched,omitempty"`
	//! IsWatched bool            `json:"watched" bson:"isWatched"`

}
