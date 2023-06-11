package model

type Message struct {
	Id        string `bson:"_id"`
	Text      string `bson:"text"`
	CreatedAt int64  `bson:"created_at"`
	StoredAt  int64  `bson:"stored_at"`
}
