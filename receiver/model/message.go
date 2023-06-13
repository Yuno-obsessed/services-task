package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	Id        primitive.ObjectID `bson:"_id"`
	Text      string             `bson:"text"`
	CreatedAt time.Time          `bson:"created_at"`
	StoredAt  time.Time          `bson:"stored_at"`
}
