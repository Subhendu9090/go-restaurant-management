package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id         primitive.ObjectID `bson:"_id"`
	Text       string             `json:"text"`
	Title      string             `json:"title"`
	Created_at time.Time          `json:"created_at"`
	Upadted_at time.Time          `json:"updated_at"`
	Note_id    string             `json:"note_id"`
}
